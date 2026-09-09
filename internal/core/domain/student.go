package domain

type Student struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
}

type CreateStudentRequest struct {
	Name     string `json:"name" validate:"required"`
	ParentID int64  `json:"parent_id" validate:"required"`
}
