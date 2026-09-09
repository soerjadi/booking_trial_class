package domain

import "time"

type PaymentAttemptStatus string

const (
	PaymentAttemptStatusPending  PaymentAttemptStatus = "pending"
	PaymentAttemptStatusSuccess  PaymentAttemptStatus = "success"
	PaymentAttemptStatusFailed   PaymentAttemptStatus = "failed"
	PaymentAttemptStatusRefunded PaymentAttemptStatus = "refunded"
)

type PaymentAttempt struct {
	ID        int64                `json:"id"`
	BookingID int64                `json:"booking_id"`
	Status    PaymentAttemptStatus `json:"status"`
	Note      string               `json:"note"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
