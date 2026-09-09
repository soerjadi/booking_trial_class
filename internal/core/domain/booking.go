package domain

import "time"

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID             int64         `json:"id"`
	TrialClassID   int64         `json:"trial_classes_id"`
	StudentID      int64         `json:"student_id"`
	Status         BookingStatus `json:"status"`
	IdempotencyKey string        `json:"-"`
	HoldExpiredAt  time.Time     `json:"-"`
	PaymentCode    string        `json:"payment_code"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type BookClassRequest struct {
	TrialClassID   int64
	StudentID      int64
	IdempotencyKey string
}
