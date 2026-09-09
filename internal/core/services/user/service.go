package user

import (
	"context"
	"fmt"

	"github.com/soerjadi/booking/internal/core/domain"
	svcport "github.com/soerjadi/booking/internal/core/ports/services"
)

type dummyService struct {
	users []domain.User
}

func NewDummyService() svcport.UserService {
	return &dummyService{
		users: []domain.User{
			{ID: 1, Username: "alice", Email: "alice@example.com", Fullname: "Alice Smith"},
			{ID: 2, Username: "bob", Email: "bob@example.com", Fullname: "Bob Jones"},
		},
	}
}

func (s *dummyService) GetAll(_ context.Context) ([]domain.User, error) {
	return s.users, nil
}

func (s *dummyService) GetByID(_ context.Context, id int64) (*domain.User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user %d not found", id)
}
