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

func TestPaymentAttemptRepository_GetByBookingID(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewPaymentAttemptRepository(mockDB)
	ctx := context.Background()

	now := time.Now()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), int64(1)).Return(mockRow{
			args: []any{int64(1), int64(1), domain.PaymentAttemptStatusPending, "", now, now},
		})

		res, err := repo.GetByBookingID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), int64(2)).Return(mockRow{
			err: errors.New("not found"),
		})

		_, err := repo.GetByBookingID(ctx, 2)
		assert.Error(t, err)
	})
}
