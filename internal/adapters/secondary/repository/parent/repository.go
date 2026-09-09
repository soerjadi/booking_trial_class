package parent

import "github.com/soerjadi/booking/internal/core/ports/repository"

type parentRepository struct {
	db repository.DB
}

func NewParentRepository(db repository.DB) repository.ParentRepositoryInterface {
	return &parentRepository{db: db}
}
