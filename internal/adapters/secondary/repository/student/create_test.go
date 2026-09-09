package student

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestStudentRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewStudentRepository(mockDB)
	ctx := context.Background()
	req := domain.Student{Name: "Jane Doe", ParentID: 1}

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.Name, req.ParentID).Return(mockRow{
			args: []any{int64(1), "Jane Doe", int64(1)},
		})

		res, err := repo.Create(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.ID)
		assert.Equal(t, "Jane Doe", res.Name)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), req.Name, req.ParentID).Return(mockRow{
			err: errors.New("insert error"),
		})

		_, err := repo.Create(ctx, req)
		assert.Error(t, err)
	})
}
