package trialclass

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

func TestHandler_Create(t *testing.T) {
	ctrl := setupMocks(t)
	defer ctrl.Finish()

	mockSvc := mock_svc.NewMockTrialClassServiceInterface(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		reqBody := domain.CreateTrialClassRequest{Name: "Class A"}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/trial_class", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockSvc.EXPECT().Create(gomock.Any(), gomock.Any()).Return(domain.TrialClass{ID: 1, Name: "Class A", Quota: domain.CLASS_QUOTA, AvailableSlots: domain.CLASS_QUOTA}, nil)

		handler.Create(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error_invalid_body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/trial_class", bytes.NewBufferString("{invalid json}"))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_service", func(t *testing.T) {
		reqBody := domain.CreateTrialClassRequest{Name: "Class A"}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/trial_class", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockSvc.EXPECT().Create(gomock.Any(), gomock.Any()).Return(domain.TrialClass{}, errors.New("internal error"))

		handler.Create(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
