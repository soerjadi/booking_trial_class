package parent

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestParentRepository_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewParentRepository(mockDB)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), int64(1)).Return(mockRow{
			args: []any{int64(1), "Jane Doe"},
		})

		res, err := repo.GetByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
		assert.Equal(t, "Jane Doe", res.Name)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), int64(2)).Return(mockRow{
			err: errors.New("not found"),
		})

		_, err := repo.GetByID(ctx, 2)
		assert.Error(t, err)
	})
}
