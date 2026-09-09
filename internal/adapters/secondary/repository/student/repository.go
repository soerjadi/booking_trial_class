package student

import "github.com/soerjadi/booking/internal/core/ports/repository"

type studentRepository struct {
	db repository.DB
}

func NewStudentRepository(db repository.DB) repository.StudentRepositoryInterface {
	return &studentRepository{db: db}
}
