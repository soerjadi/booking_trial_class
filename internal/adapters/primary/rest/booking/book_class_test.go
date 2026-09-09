package booking

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_svc "github.com/soerjadi/booking/internal/core/ports/services/mocks"
)

func TestHandler_BookClass(t *testing.T) {
	ctrl := setupMocks(t)
	defer ctrl.Finish()

	mockSvc := mock_svc.NewMockBookingServiceInterface(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		reqBody := domain.BookClassRequest{StudentID: 1, TrialClassID: 1}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer(body))
		req.Header.Set("X-Idempotency-Key", "idemp-key-1")
		w := httptest.NewRecorder()

		reqBody.IdempotencyKey = "idemp-key-1"
		mockSvc.EXPECT().BookClass(gomock.Any(), reqBody).Return(domain.Booking{ID: 1}, nil)

		handler.BookClass(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error_missing_header", func(t *testing.T) {
		reqBody := domain.BookClassRequest{StudentID: 1, TrialClassID: 1}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.BookClass(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_invalid_body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/booking", bytes.NewBufferString("{invalid json}"))
		req.Header.Set("X-Idempotency-Key", "idemp-key-1")
		w := httptest.NewRecorder()

		handler.BookClass(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_service", func(t *testing.T) {
		reqBody := domain.BookClassRequest{StudentID: 1, TrialClassID: 1}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer(body))
		req.Header.Set("X-Idempotency-Key", "idemp-key-1")
		w := httptest.NewRecorder()

		reqBody.IdempotencyKey = "idemp-key-1"
		mockSvc.EXPECT().BookClass(gomock.Any(), reqBody).Return(domain.Booking{}, errors.New("service error"))

		handler.BookClass(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
