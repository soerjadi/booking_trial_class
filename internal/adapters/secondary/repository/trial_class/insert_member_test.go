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

func TestTrialClassRepository_InsertMember(t *testing.T) {
	ctrl, _ := setupMocks(t)
	defer ctrl.Finish()

	mockDB := mock_repo.NewMockDB(ctrl)
	repo := NewTrialClassRepository(mockDB)
	ctx := context.Background()

	req := domain.TrialClassMember{TrialClassID: 1, StudentID: 1}

	t.Run("success", func(t *testing.T) {
		mockDB.EXPECT().Exec(ctx, gomock.Any(), req.TrialClassID, req.StudentID).Return(pgconn.CommandTag{}, nil)

		err := repo.InsertMember(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockDB.EXPECT().Exec(ctx, gomock.Any(), req.TrialClassID, req.StudentID).Return(pgconn.CommandTag{}, errors.New("insert error"))

		err := repo.InsertMember(ctx, req)
		assert.Error(t, err)
	})
}
