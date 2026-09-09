package trial_class

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestTrialClassRepository_Create(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewTrialClassRepository(mockDB)
	ctx := context.Background()

	req := domain.TrialClass{Name: "Class A", Quota: int64(10), AvailableSlots: int64(10)}

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.Name, req.Quota, req.AvailableSlots).Return(mockRow{
			args: []any{int64(1), req.Name, req.Quota, req.AvailableSlots},
		})

		res, err := repo.Create(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.Name, req.Quota, req.AvailableSlots).Return(mockRow{
			err: errors.New("insert error"),
		})

		_, err := repo.Create(ctx, req)
		assert.Error(t, err)
	})
}
