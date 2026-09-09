package student

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
	svc svcport.StudentServiceInterface
}

func NewHandler(svc svcport.StudentServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Mount(r *mux.Router) {
	r.HandleFunc("/students", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/students/{id:[0-9]+}", h.GetByID).Methods(http.MethodGet)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateStudentRequest
	if err := json.ConfigDefault.NewDecoder(r.Body).Decode(&request); err != nil {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: err.Error()})
		return
	}

	student, err := h.svc.Create(r.Context(), domain.Student{
		Name:     request.Name,
		ParentID: request.ParentID,
	})
	if err != nil {
		rest.WriteJSON(w, http.StatusInternalServerError, rest.ErrorResponse{Message: err.Error()})
		return
	}
	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "success", Data: student})
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		rest.WriteJSON(w, http.StatusBadRequest, rest.ErrorResponse{Message: "invalid id"})
		return
	}

	student, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		rest.WriteJSON(w, http.StatusInternalServerError, rest.ErrorResponse{Message: err.Error()})
		return
	}
	rest.WriteJSON(w, http.StatusOK, rest.SuccessResponse{Message: "success", Data: student})
}
