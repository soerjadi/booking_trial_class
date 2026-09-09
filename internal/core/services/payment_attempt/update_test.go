package paymentattempt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soerjadi/booking/internal/core/domain"
)

func TestPaymentAttemptService_Update(t *testing.T) {
	ctrl, mockRepo, mockBookingRepo, mockClassRepo := setupMocks(t)
	defer ctrl.Finish()

	service := NewPaymentAttemptService(mockRepo, mockBookingRepo, mockClassRepo)

	ctx := context.Background()
	req := domain.PaymentAttempt{ID: 1, BookingID: 1}

	mockRepo.EXPECT().Update(ctx, req).Return(nil)

	err := service.Update(ctx, req)
	assert.NoError(t, err)
}
