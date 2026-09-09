package student

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/soerjadi/booking/internal/core/domain"
	mock_svc "github.com/soerjadi/booking/internal/core/ports/services/mocks"
)

func TestHandler_GetByID(t *testing.T) {
	ctrl := setupMocks(t)
	defer ctrl.Finish()

	mockSvc := mock_svc.NewMockStudentServiceInterface(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/students/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		mockSvc.EXPECT().GetByID(gomock.Any(), int64(1)).Return(domain.Student{ID: 1, Name: "John"}, nil)

		handler.GetByID(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error_invalid_id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/students/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})
		w := httptest.NewRecorder()

		handler.GetByID(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_service", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/students/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		mockSvc.EXPECT().GetByID(gomock.Any(), int64(1)).Return(domain.Student{}, errors.New("not found"))

		handler.GetByID(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
