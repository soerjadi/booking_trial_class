package student

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_repo "github.com/soerjadi/booking/internal/core/ports/repository/mocks"
)

func TestStudentService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockStudentRepositoryInterface(ctrl)
	service := NewStudentService(mockRepo)

	ctx := context.Background()
	req := domain.Student{Name: "John Doe"}
	resp := domain.Student{ID: 1, Name: "John Doe"}

	mockRepo.EXPECT().Create(ctx, req).Return(resp, nil)

	res, err := service.Create(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, resp, res)
}
