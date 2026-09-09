package payment

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_svc "github.com/soerjadi/booking/internal/core/ports/services/mocks"
)

func TestHandler_Settlement(t *testing.T) {
	ctrl := setupMocks(t)
	defer ctrl.Finish()

	mockSvc := mock_svc.NewMockPaymentAttemptServiceInterface(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		reqBody := SettlementRequest{PaymentCode: "PAY-123"}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/payment/settlement", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockSvc.EXPECT().Settlement(gomock.Any(), "PAY-123").Return(nil)

		handler.Settlement(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error_invalid_body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/payment/settlement", bytes.NewBufferString("{invalid json}"))
		w := httptest.NewRecorder()

		handler.Settlement(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_empty_payment_code", func(t *testing.T) {
		reqBody := SettlementRequest{PaymentCode: ""}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/payment/settlement", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.Settlement(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_service", func(t *testing.T) {
		reqBody := SettlementRequest{PaymentCode: "PAY-123"}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/payment/settlement", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockSvc.EXPECT().Settlement(gomock.Any(), "PAY-123").Return(errors.New("settlement error"))

		handler.Settlement(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
