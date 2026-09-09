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

func TestBookingRepository_GetByStudentID(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewBookingRepository(mockDB)
	ctx := context.Background()
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().Query(ctx, gomock.Any(), int64(1)).Return(&mockRows{
			rows: [][]any{
				{int64(1), int64(1), int64(1), domain.BookingStatusPending, "test-key", now, "PAY-123", now, now},
			},
		}, nil)

		res, err := repo.GetByStudentID(ctx, 1)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, int64(1), res[0].ID)
	})

	t.Run("error_query", func(t *testing.T) {
		mockDB.EXPECT().Query(ctx, gomock.Any(), int64(2)).Return(nil, errors.New("db error"))

		_, err := repo.GetByStudentID(ctx, 2)
		assert.Error(t, err)
	})
}
