package trial_class

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestTrialClassRepository_GetMember(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewTrialClassRepository(mockDB)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().Query(ctx, gomock.Any(), int64(1)).Return(&mockRows{
			rows: [][]any{
				{int64(1), int64(1), int64(1)},
			},
		}, nil)

		res, err := repo.GetMember(ctx, 1)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, int64(1), res[0].ID)
	})

	t.Run("error_query", func(t *testing.T) {
		mockDB.EXPECT().Query(ctx, gomock.Any(), int64(2)).Return(nil, errors.New("db error"))

		_, err := repo.GetMember(ctx, 2)
		assert.Error(t, err)
	})
}
