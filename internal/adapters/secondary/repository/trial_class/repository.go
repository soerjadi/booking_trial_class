package trial_class

import "github.com/soerjadi/booking/internal/core/ports/repository"

type trialClassRepository struct {
	db repository.DB
}

func NewTrialClassRepository(db repository.DB) repository.TrialClassRepositoryInterface {
	return &trialClassRepository{db: db}
}
