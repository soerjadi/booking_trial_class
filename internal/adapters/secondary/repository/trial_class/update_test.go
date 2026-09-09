package trial_class

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestTrialClassRepository_Update(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewTrialClassRepository(mockDB)
	ctx := context.Background()

	req := domain.TrialClass{ID: 1, Name: "Class A", Quota: int64(10), AvailableSlots: int64(9)}

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().Exec(ctx, gomock.Any(), req.Name, req.Quota, req.AvailableSlots, req.ID).Return(pgconn.CommandTag{}, nil)

		err := repo.Update(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().Exec(ctx, gomock.Any(), req.Name, req.Quota, req.AvailableSlots, req.ID).Return(pgconn.CommandTag{}, errors.New("update error"))

		err := repo.Update(ctx, req)
		assert.Error(t, err)
	})
}
