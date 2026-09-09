package booking

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
)

func TestBookingService_BookClass(t *testing.T) {
	ctrl, mockRepo, mockClassRepo, mockStudentRepo, mockPaymentRepo := setupMocks(t)
	defer ctrl.Finish()

	service := NewBookingService(mockRepo, mockClassRepo, mockStudentRepo, mockPaymentRepo)

	ctx := context.Background()
	req := domain.BookClassRequest{TrialClassID: 1, StudentID: 1, IdempotencyKey: "test-key"}

	t.Run("success", func(t *testing.T) {
		trialClass := domain.TrialClass{ID: 1, AvailableSlots: 5}
		student := domain.Student{ID: 1}
		members := []domain.TrialClassMember{}

		mockClassRepo.EXPECT().GetByID(gomock.Any(), req.TrialClassID).Return(trialClass, nil)
		mockStudentRepo.EXPECT().GetByID(gomock.Any(), req.StudentID).Return(student, nil)
		mockClassRepo.EXPECT().GetMember(gomock.Any(), req.TrialClassID).Return(members, nil)

		// Inside db.Do transaction
		mockClassRepo.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, tc domain.TrialClass) error {
			assert.Equal(t, int64(4), tc.AvailableSlots)
			return nil
		})

		mockRepo.EXPECT().Create(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, b domain.Booking) (domain.Booking, error) {
			assert.Equal(t, domain.BookingStatusPending, b.Status)
			assert.NotEmpty(t, b.PaymentCode)
			b.ID = int64(1)
			return b, nil
		})

		mockPaymentRepo.EXPECT().Create(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, p domain.PaymentAttempt) (domain.PaymentAttempt, error) {
			assert.Equal(t, int64(1), p.BookingID)
			assert.Equal(t, domain.PaymentAttemptStatusPending, p.Status)
			return p, nil
		})

		res, err := service.BookClass(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
		assert.Equal(t, domain.BookingStatusPending, res.Status)
	})

	t.Run("class is full", func(t *testing.T) {
		trialClass := domain.TrialClass{ID: 1, AvailableSlots: 0}
		student := domain.Student{ID: 1}
		members := []domain.TrialClassMember{}

		mockClassRepo.EXPECT().GetByID(gomock.Any(), req.TrialClassID).Return(trialClass, nil)
		mockStudentRepo.EXPECT().GetByID(gomock.Any(), req.StudentID).Return(student, nil)
		mockClassRepo.EXPECT().GetMember(gomock.Any(), req.TrialClassID).Return(members, nil)

		_, err := service.BookClass(ctx, req)
		assert.EqualError(t, err, "trial class is full")
	})

	t.Run("student already registered", func(t *testing.T) {
		trialClass := domain.TrialClass{ID: 1, AvailableSlots: 5}
		student := domain.Student{ID: 1}
		members := []domain.TrialClassMember{{StudentID: 1}}

		mockClassRepo.EXPECT().GetByID(gomock.Any(), req.TrialClassID).Return(trialClass, nil)
		mockStudentRepo.EXPECT().GetByID(gomock.Any(), req.StudentID).Return(student, nil)
		mockClassRepo.EXPECT().GetMember(gomock.Any(), req.TrialClassID).Return(members, nil)

		_, err := service.BookClass(ctx, req)
		assert.EqualError(t, err, "student already register in this class")
	})

	t.Run("error getting entity", func(t *testing.T) {
		mockClassRepo.EXPECT().GetByID(gomock.Any(), req.TrialClassID).Return(domain.TrialClass{}, errors.New("db error"))
		mockStudentRepo.EXPECT().GetByID(gomock.Any(), req.StudentID).Return(domain.Student{}, nil)
		mockClassRepo.EXPECT().GetMember(gomock.Any(), req.TrialClassID).Return([]domain.TrialClassMember{}, nil)

		_, err := service.BookClass(ctx, req)
		assert.EqualError(t, err, "db error")
	})
}
