package payment

import (
	"net/http"

	json "github.com/bytedance/sonic"
	"github.com/gorilla/mux"
	"github.com/soerjadi/booking/internal/adapters/primary/rest"
	svcport "github.com/soerjadi/booking/internal/core/ports/services"
)

type Handler struct {
	svc svcport.PaymentAttemptServiceInterface
}

func NewHandler(svc svcport.PaymentAttemptServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Mount(r *mux.Router) {
	r.HandleFunc("/payment/settlement", h.Settlement).Methods(http.MethodPost)
}

type SettlementRequest struct {
	PaymentCode string `json:"payment_code"`
}

func (h *Handler) Settlement(w http.ResponseWriter, r *http.Request) {
	var request SettlementRequest
	if err := json.ConfigDefault.NewDecoder(r.Body).Decode(&request); err != nil {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: err.Error()})
		return
	}

	if request.PaymentCode == "" {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: "payment_code is required"})
		return
	}

	err := h.svc.Settlement(r.Context(), request.PaymentCode)
	if err != nil {
		rest.WriteJSON(w, http.StatusInternalServerError, rest.ErrorResponse{Message: err.Error()})
		return
	}

	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "payment confirmed", Data: nil})
}
