package trialclass

import (
	"net/http"

	json "github.com/bytedance/sonic"
	"github.com/gorilla/mux"
	"github.com/soerjadi/booking/internal/adapters/primary/rest"
	"github.com/soerjadi/booking/internal/core/domain"
	svcport "github.com/soerjadi/booking/internal/core/ports/services"
)

type Handler struct {
	svc svcport.TrialClassServiceInterface
}

func NewHandler(svc svcport.TrialClassServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Mount(r *mux.Router) {
	r.HandleFunc("/trial_class", h.Create).Methods(http.MethodPost)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateTrialClassRequest
	if err := json.ConfigDefault.NewDecoder(r.Body).Decode(&request); err != nil {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: err.Error()})
		return
	}

	trialClass, err := h.svc.Create(r.Context(), domain.TrialClass{
		Name:           request.Name,
		Quota:          domain.CLASS_QUOTA,
		AvailableSlots: domain.CLASS_QUOTA,
	})
	if err != nil {
		rest.WriteJSON(w, http.StatusInternalServerError, rest.ErrorResponse{Message: err.Error()})
		return
	}
	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "success", Data: trialClass})
}
