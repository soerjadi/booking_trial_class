package booking

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

func TestHandler_GetByStudentID(t *testing.T) {
	ctrl := setupMocks(t)
	defer ctrl.Finish()

	mockSvc := mock_svc.NewMockBookingServiceInterface(ctrl)
	handler := NewHandler(mockSvc)

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/students/1/bookings", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		mockSvc.EXPECT().GetByStudentID(gomock.Any(), int64(1)).Return([]domain.Booking{{ID: 1}}, nil)

		handler.GetByStudentID(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error_invalid_id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/students/abc/bookings", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})
		w := httptest.NewRecorder()

		handler.GetByStudentID(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error_service", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/students/1/bookings", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		mockSvc.EXPECT().GetByStudentID(gomock.Any(), int64(1)).Return(nil, errors.New("service error"))

		handler.GetByStudentID(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
