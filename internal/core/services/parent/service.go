package parent

import (
	"github.com/soerjadi/booking/internal/core/ports/repository"
	"github.com/soerjadi/booking/internal/core/ports/services"
)

type parentService struct {
	repo repository.ParentRepositoryInterface
}

func NewParentService(repo repository.ParentRepositoryInterface) services.ParentServiceInterface {
	return &parentService{repo: repo}
}
