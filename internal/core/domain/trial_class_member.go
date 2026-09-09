package domain

type TrialClassMember struct {
	ID           int64 `json:"id"`
	TrialClassID int64 `json:"trial_classes_id"`
	StudentID    int64 `json:"student_id"`
}
