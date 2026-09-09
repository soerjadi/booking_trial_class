package paymentattempt

import (
	"context"
	"errors"
	"time"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (s *paymentAttemptService) Settlement(ctx context.Context, paymentCode string) error {
	booking, err := s.repoBooking.GetByPaymentCode(ctx, paymentCode)
	if err != nil {
		log.ErrorCtx(ctx, "[service.paymentattempt.Settlement.GetByPaymentCode] failed get booking", log.Field("paymentCode", paymentCode), log.Field("error", err))
		return err
	}

	trialClass, err := s.repoClass.GetByID(ctx, booking.TrialClassID)
	if err != nil {
		log.ErrorCtx(ctx, "[service.paymentattempt.Settlement.GetByID] failed get trial class", log.Field("booking", booking), log.Field("error", err))
		return err
	}

	rollbackSlot := func(ctx context.Context) error {
		return db.Do(ctx, func(ctx context.Context) error {
			updateClassReq := domain.TrialClass{
				ID:             trialClass.ID,
				Name:           trialClass.Name,
				Quota:          trialClass.Quota,
				AvailableSlots: trialClass.AvailableSlots + 1,
			}
			if err := s.repoClass.Update(ctx, updateClassReq); err != nil {
				return err
			}
			booking.Status = domain.BookingStatusCancelled
			return s.repoBooking.Update(ctx, booking)
		})
	}

	// if already expired then it should be update trial_class.available_slots + 1
	if time.Now().UTC().After(booking.HoldExpiredAt) {
		if err := rollbackSlot(ctx); err != nil {
			log.ErrorCtx(ctx, "[service.paymentattempt.Settlement] failed rollback expired booking", log.Field("booking", booking), log.Field("error", err))
		}
		return errors.New("booking hold expired, slot released")
	}

	if booking.Status != domain.BookingStatusPending {
		return errors.New("booking is not pending")
	}

	paymentAttempt, err := s.repo.GetByBookingID(ctx, booking.ID)
	if err != nil {
		log.ErrorCtx(ctx, "[service.paymentattempt.Settlement.GetByBookingID] failed get payment attempt", log.Field("booking", booking), log.Field("error", err))
		return err
	}

	err = db.Do(ctx, func(ctx context.Context) error {
		// Update booking status to success and clear HoldExpiredAt
		booking.Status = domain.BookingStatusConfirmed
		booking.HoldExpiredAt = time.Time{}
		if err := s.repoBooking.Update(ctx, booking); err != nil {
			return err
		}

		// Update payment attempt status to success
		paymentAttempt.Status = domain.PaymentAttemptStatusSuccess
		if err := s.repo.Update(ctx, paymentAttempt); err != nil {
			return err
		}

		// also add trial_member after all success.
		member := domain.TrialClassMember{
			TrialClassID: booking.TrialClassID,
			StudentID:    booking.StudentID,
		}
		if err := s.repoClass.InsertMember(ctx, member); err != nil {
			return err
		}

		return nil
	})

	// if it fail then don forget to rollback available_slots
	if err != nil {
		log.ErrorCtx(ctx, "[service.paymentattempt.Settlement] failed settlement transaction, rolling back slots", log.Field("booking", booking), log.Field("error", err))
		_ = rollbackSlot(ctx)
		return err
	}

	return nil
}
