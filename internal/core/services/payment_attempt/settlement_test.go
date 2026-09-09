package paymentattempt

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
)

func TestPaymentAttemptService_Settlement(t *testing.T) {
	ctrl, mockRepo, mockBookingRepo, mockClassRepo := setupMocks(t)
	defer ctrl.Finish()

	service := NewPaymentAttemptService(mockRepo, mockBookingRepo, mockClassRepo)
	ctx := context.Background()
	paymentCode := "PAY-123"

	t.Run("success", func(t *testing.T) {
		booking := domain.Booking{ID: 1, TrialClassID: 1, StudentID: 1, Status: domain.BookingStatusPending, HoldExpiredAt: time.Now().Add(1 * time.Hour)}
		trialClass := domain.TrialClass{ID: 1, AvailableSlots: 5}
		paymentAttempt := domain.PaymentAttempt{ID: 1, BookingID: 1, Status: domain.PaymentAttemptStatusPending}

		mockBookingRepo.EXPECT().GetByPaymentCode(ctx, paymentCode).Return(booking, nil)
		mockClassRepo.EXPECT().GetByID(ctx, booking.TrialClassID).Return(trialClass, nil)
		mockRepo.EXPECT().GetByBookingID(ctx, booking.ID).Return(paymentAttempt, nil)

		// Inside transaction
		mockBookingRepo.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, b domain.Booking) error {
			assert.Equal(t, domain.BookingStatusConfirmed, b.Status)
			assert.True(t, b.HoldExpiredAt.IsZero())
			return nil
		})
		mockRepo.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, p domain.PaymentAttempt) error {
			assert.Equal(t, domain.PaymentAttemptStatusSuccess, p.Status)
			return nil
		})
		mockClassRepo.EXPECT().InsertMember(ctx, domain.TrialClassMember{TrialClassID: 1, StudentID: 1}).Return(nil)

		err := service.Settlement(ctx, paymentCode)
		assert.NoError(t, err)
	})

	t.Run("booking not found", func(t *testing.T) {
		mockBookingRepo.EXPECT().GetByPaymentCode(ctx, paymentCode).Return(domain.Booking{}, errors.New("not found"))
		err := service.Settlement(ctx, paymentCode)
		assert.Error(t, err)
	})

	t.Run("expired hold", func(t *testing.T) {
		booking := domain.Booking{ID: 1, TrialClassID: 1, Status: domain.BookingStatusPending, HoldExpiredAt: time.Now().Add(-1 * time.Hour)}
		trialClass := domain.TrialClass{ID: 1, AvailableSlots: 5}
		paymentAttempt := domain.PaymentAttempt{ID: 1, BookingID: 1, Status: domain.PaymentAttemptStatusPending}

		mockBookingRepo.EXPECT().GetByPaymentCode(ctx, paymentCode).Return(booking, nil)
		mockClassRepo.EXPECT().GetByID(ctx, booking.TrialClassID).Return(trialClass, nil)
		mockRepo.EXPECT().GetByBookingID(ctx, booking.ID).Return(paymentAttempt, nil)

		// Rollback logic
		mockClassRepo.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, tc domain.TrialClass) error {
			assert.Equal(t, int64(6), tc.AvailableSlots)
			return nil
		})
		mockRepo.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, p domain.PaymentAttempt) error {
			assert.Equal(t, domain.PaymentAttemptStatusRefunded, p.Status)
			return nil
		})
		mockBookingRepo.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, b domain.Booking) error {
			assert.Equal(t, domain.BookingStatusCancelled, b.Status)
			return nil
		})

		err := service.Settlement(ctx, paymentCode)
		assert.EqualError(t, err, "booking hold expired, slot released")
	})

	t.Run("not pending", func(t *testing.T) {
		booking := domain.Booking{ID: 1, TrialClassID: 1, Status: domain.BookingStatusConfirmed, HoldExpiredAt: time.Now().Add(1 * time.Hour)}
		trialClass := domain.TrialClass{ID: 1, AvailableSlots: 5}
		paymentAttempt := domain.PaymentAttempt{ID: 1, BookingID: 1, Status: domain.PaymentAttemptStatusSuccess}

		mockBookingRepo.EXPECT().GetByPaymentCode(ctx, paymentCode).Return(booking, nil)
		mockClassRepo.EXPECT().GetByID(ctx, booking.TrialClassID).Return(trialClass, nil)
		mockRepo.EXPECT().GetByBookingID(ctx, booking.ID).Return(paymentAttempt, nil)

		err := service.Settlement(ctx, paymentCode)
		assert.EqualError(t, err, "booking is not pending")
	})
}
