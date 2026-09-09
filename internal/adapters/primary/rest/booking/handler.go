package booking

import (
	"net/http"
	"strconv"

	json "github.com/bytedance/sonic"
	"github.com/gorilla/mux"
	"github.com/soerjadi/booking/internal/adapters/primary/rest"
	"github.com/soerjadi/booking/internal/core/domain"
	svcport "github.com/soerjadi/booking/internal/core/ports/services"
)

type Handler struct {
	svc svcport.BookingServiceInterface
}

func NewHandler(svc svcport.BookingServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Mount(r *mux.Router) {
	r.HandleFunc("/booking", h.BookClass).Methods(http.MethodPost)
	r.HandleFunc("/students/{id:[0-9]+}/bookings", h.GetByStudentID).Methods(http.MethodGet)
}

func (h *Handler) BookClass(w http.ResponseWriter, r *http.Request) {
	idempotencyKey := r.Header.Get("X-Idempotency-Key")
	if idempotencyKey == "" {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: "X-Idempotency-Key header is required"})
		return
	}

	var request domain.BookClassRequest
	if err := json.ConfigDefault.NewDecoder(r.Body).Decode(&request); err != nil {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: err.Error()})
		return
	}

	request.IdempotencyKey = idempotencyKey

	booking, err := h.svc.BookClass(r.Context(), request)
	if err != nil {
		rest.WriteJSON(w, http.StatusInternalServerError, rest.ErrorResponse{Message: err.Error()})
		return
	}
	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "success", Data: booking})
}

func (h *Handler) GetByStudentID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: "invalid id"})
		return
	}

	bookings, err := h.svc.GetByStudentID(r.Context(), id)
	if err != nil {
		rest.WriteJSON(w, http.StatusInternalServerError, rest.ErrorResponse{Message: err.Error()})
		return
	}
	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "success", Data: bookings})
}
