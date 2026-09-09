package domain

type Parent struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateParentRequest struct {
	Name string `json:"name" validate:"required"`
}
