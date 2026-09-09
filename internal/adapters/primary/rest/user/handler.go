package user

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soerjadi/booking/internal/adapters/primary/rest"
	svcport "github.com/soerjadi/booking/internal/core/ports/services"
)

type Handler struct {
	svc svcport.UserService
}

func NewHandler(svc svcport.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Mount(r *mux.Router) {
	r.HandleFunc("/users", h.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", h.GetByID).Methods(http.MethodGet)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.GetAll(r.Context())
	if err != nil {
		rest.WriteJSON(w, http.StatusInternalServerError, rest.ErrorResponse{Message: err.Error()})
		return
	}
	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "success", Data: users})
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	u, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		rest.WriteJSON(w, http.StatusNotFound, rest.ErrorResponse{Message: err.Error()})
		return
	}
	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "success", Data: u})
}
