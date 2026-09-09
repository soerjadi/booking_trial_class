package payment_attempt

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

func TestPaymentAttemptRepository_Create(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewPaymentAttemptRepository(mockDB)
	ctx := context.Background()

	req := domain.PaymentAttempt{BookingID: 1, Status: domain.PaymentAttemptStatusPending}
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.BookingID, req.Status, req.Note).Return(mockRow{
			args: []any{int64(1), int64(1), domain.PaymentAttemptStatusPending, "", now, now},
		})

		res, err := repo.Create(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.BookingID, req.Status, req.Note).Return(mockRow{
			err: errors.New("insert error"),
		})

		_, err := repo.Create(ctx, req)
		assert.Error(t, err)
	})
}
