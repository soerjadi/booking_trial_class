package paymentattempt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soerjadi/booking/internal/core/domain"
)

func TestPaymentAttemptService_Create(t *testing.T) {
	ctrl, mockRepo, mockBookingRepo, mockClassRepo := setupMocks(t)
	defer ctrl.Finish()

	service := NewPaymentAttemptService(mockRepo, mockBookingRepo, mockClassRepo)

	ctx := context.Background()
	req := domain.PaymentAttempt{BookingID: 1}
	resp := domain.PaymentAttempt{ID: 1, BookingID: 1}

	mockRepo.EXPECT().Create(ctx, req).Return(resp, nil)

	res, err := service.Create(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, resp, res)
}
