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

func TestBookingRepository_GetByPaymentCode(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewBookingRepository(mockDB)
	ctx := context.Background()
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), "PAY-123").Return(mockRow{
			args: []any{int64(1), int64(1), int64(1), domain.BookingStatusPending, "test-key", now, "PAY-123", now, now},
		})

		res, err := repo.GetByPaymentCode(ctx, "PAY-123")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), "PAY-123").Return(mockRow{
			err: errors.New("not found"),
		})

		_, err := repo.GetByPaymentCode(ctx, "PAY-123")
		assert.Error(t, err)
	})
}
