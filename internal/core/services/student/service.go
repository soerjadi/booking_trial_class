package student

import (
	"github.com/soerjadi/booking/internal/core/ports/repository"
	"github.com/soerjadi/booking/internal/core/ports/services"
)

type studentService struct {
	repo repository.StudentRepositoryInterface
}

func NewStudentService(repo repository.StudentRepositoryInterface) services.StudentServiceInterface {
	return &studentService{repo: repo}
}
