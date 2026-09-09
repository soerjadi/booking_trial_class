package booking

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soerjadi/booking/internal/core/domain"
)

func TestBookingService_Update(t *testing.T) {
	ctrl, mockRepo, mockClassRepo, mockStudentRepo, mockPaymentRepo := setupMocks(t)
	defer ctrl.Finish()

	service := NewBookingService(mockRepo, mockClassRepo, mockStudentRepo, mockPaymentRepo)

	ctx := context.Background()
	req := domain.Booking{ID: 1, Status: domain.BookingStatusConfirmed}

	mockRepo.EXPECT().Update(ctx, req).Return(nil)

	err := service.Update(ctx, req)
	assert.NoError(t, err)
}
