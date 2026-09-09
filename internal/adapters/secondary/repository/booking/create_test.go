package booking

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestBookingRepository_Create(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewBookingRepository(mockDB)
	ctx := context.Background()

	req := domain.Booking{TrialClassID: 1, StudentID: 1, PaymentCode: "PAY-123", Status: domain.BookingStatusPending, IdempotencyKey: "test-key"}
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.TrialClassID, req.StudentID, req.Status, req.IdempotencyKey, req.HoldExpiredAt, req.PaymentCode).Return(mockRow{
			args: []any{int64(1), req.TrialClassID, req.StudentID, req.Status, req.IdempotencyKey, req.HoldExpiredAt, req.PaymentCode, now, now},
		})

		res, err := repo.Create(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.TrialClassID, req.StudentID, req.Status, req.IdempotencyKey, req.HoldExpiredAt, req.PaymentCode).Return(mockRow{
			err: errors.New("insert error"),
		})

		_, err := repo.Create(ctx, req)
		assert.Error(t, err)
	})
}
