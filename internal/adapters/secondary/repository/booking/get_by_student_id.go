package booking

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *bookingRepository) GetByStudentID(ctx context.Context, studentID int64) ([]domain.Booking, error) {
	var bookings []domain.Booking
	query := `
		SELECT id, trial_classes_id, student_id, status, idempotency_key, hold_expires_at, payment_code, created_at, updated_at 
		FROM bookings 
		WHERE student_id = $1
	`
	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b domain.Booking
		if err := rows.Scan(&b.ID, &b.TrialClassID, &b.StudentID, &b.Status, &b.IdempotencyKey, &b.HoldExpiredAt, &b.PaymentCode, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}
