package parent

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

	mockSvc := mock_svc.NewMockParentServiceInterface(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		reqBody := domain.CreateParentRequest{Name: "Jane"}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/parents", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockSvc.EXPECT().Create(gomock.Any(), gomock.Any()).Return(domain.Parent{ID: 1, Name: "Jane"}, nil)

		handler.Create(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error_invalid_body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/parents", bytes.NewBufferString("{invalid json}"))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_service", func(t *testing.T) {
		reqBody := domain.CreateParentRequest{Name: "Jane"}
		body, _ := sonic.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/parents", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockSvc.EXPECT().Create(gomock.Any(), gomock.Any()).Return(domain.Parent{}, errors.New("internal error"))

		handler.Create(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
