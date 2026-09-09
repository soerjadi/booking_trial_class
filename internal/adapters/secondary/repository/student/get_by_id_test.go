package student

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestStudentRepository_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewStudentRepository(mockDB)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), int64(1)).Return(mockRow{
			args: []any{int64(1), "John Doe", int64(1)},
		})

		student, err := repo.GetByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), student.ID)
		assert.Equal(t, "John Doe", student.Name)
		assert.Equal(t, int64(1), student.ParentID)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), int64(2)).Return(mockRow{
			err: errors.New("not found"),
		})

		_, err := repo.GetByID(ctx, 2)
		assert.Error(t, err)
	})
}
