package trialclass

import (
	"github.com/soerjadi/booking/internal/core/ports/repository"
	"github.com/soerjadi/booking/internal/core/ports/services"
)

type trialClassService struct {
	repo repository.TrialClassRepositoryInterface
}

func NewTrialClassService(repo repository.TrialClassRepositoryInterface) services.TrialClassServiceInterface {
	return &trialClassService{repo: repo}
}
