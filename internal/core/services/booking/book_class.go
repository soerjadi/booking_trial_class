package booking

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
	"golang.org/x/sync/errgroup"
)

func (s *bookingService) BookClass(ctx context.Context, request domain.BookClassRequest) (domain.Booking, error) {
	var (
		trialClass   domain.TrialClass
		student      domain.Student
		classMembers []domain.TrialClassMember
		booking      domain.Booking
	)

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		trialClass, err = s.repoClass.GetByID(gCtx, request.TrialClassID)
		if err != nil {
			log.ErrorCtx(gCtx, "[service.booking.BookClass.GetByID] failed get trial class by id", log.Field("request", request.TrialClassID), log.Field("error", err))
		}
		return err
	})

	g.Go(func() error {
		var err error
		student, err = s.repoStudent.GetByID(gCtx, request.StudentID)
		if err != nil {
			log.ErrorCtx(gCtx, "[service.booking.BookClass.GetByID] failed get student by id", log.Field("request", request.StudentID), log.Field("error", err))
		}
		return err
	})

	g.Go(func() error {
		var err error
		classMembers, err = s.repoClass.GetMember(gCtx, request.TrialClassID)
		if err != nil {
			log.ErrorCtx(gCtx, "[service.booking.BookClass.GetMember] failed get class members", log.Field("trialClassID", request.TrialClassID), log.Field("error", err))
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.ErrorCtx(ctx, "[service.booking.BookClass] failed to get trial class, student, or class members", log.Field("request", request), log.Field("error", err))
		return domain.Booking{}, err
	}

	// check available slots. does it full or not
	if trialClass.AvailableSlots <= 0 {
		log.ErrorCtx(ctx, "[service.booking.BookClass.GetByID] trial class is full", log.Field("request", request.TrialClassID), log.Field("trialClass", trialClass))
		return domain.Booking{}, errors.New("trial class is full")
	}

	// check if student already register in this class
	// *since this was only capped 4, using iteration wont consume much memory
	for _, member := range classMembers {
		if member.StudentID == student.ID {
			return domain.Booking{}, errors.New("student already register in this class")
		}
	}

	// do transaction to update available slot and create booking
	err := db.Do(ctx, func(ctx context.Context) error {
		// do update class to reduce available slot - 1
		updateClassReq := domain.TrialClass{
			ID:             trialClass.ID,
			Name:           trialClass.Name,
			Quota:          trialClass.Quota,
			AvailableSlots: trialClass.AvailableSlots - 1,
		}
		err := s.repoClass.Update(ctx, updateClassReq)
		if err != nil {
			log.ErrorCtx(ctx, "[service.booking.BookClass.Update] failed update trial class", log.Field("request", updateClassReq), log.Field("error", err))
			return err
		}

		// create booking with status pending
		bookRequest := domain.Booking{
			TrialClassID:   request.TrialClassID,
			StudentID:      request.StudentID,
			IdempotencyKey: request.IdempotencyKey,
			Status:         domain.BookingStatusPending,
			PaymentCode:    s.generatePaymentCode(),
			HoldExpiredAt:  time.Now().UTC().Add(5 * time.Minute),
		}

		booking, err = s.repo.Create(ctx, bookRequest)
		if err != nil {
			log.ErrorCtx(ctx, "[service.booking.BookClass.Create] failed create booking", log.Field("request", bookRequest), log.Field("error", err))
			return err
		}

		// create payment attempt with status pending
		paymentReq := domain.PaymentAttempt{
			BookingID: booking.ID,
			Status:    domain.PaymentAttemptStatusPending,
		}
		_, err = s.repoPayment.Create(ctx, paymentReq)
		if err != nil {
			log.ErrorCtx(ctx, "[service.booking.BookClass.Create] failed create payment attempt", log.Field("request", paymentReq), log.Field("error", err))
			return err
		}

		return nil
	})
	if err != nil {
		log.ErrorCtx(ctx, "[service.booking.BookClass.Update] failed update trial class", log.Field("request", request), log.Field("error", err))
		return domain.Booking{}, err
	}

	return booking, nil
}

func (s *bookingService) generatePaymentCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		code[i] = charset[n.Int64()]
	}
	return string(code)
}
