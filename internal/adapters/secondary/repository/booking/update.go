package booking

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *bookingRepository) Update(ctx context.Context, request domain.Booking) error {
	query := `
	UPDATE
		bookings
	SET
		trial_classes_id = $1,
		student_id = $2,
		status = $3,
		idempotency_key = $4,
		hold_expires_at = $5,
		payment_code = $6,
		updated_at = NOW()
	WHERE
		id = $7
	`

	_, err := r.db.Exec(ctx, query, request.TrialClassID, request.StudentID, request.Status, request.IdempotencyKey, request.HoldExpiredAt, request.PaymentCode, request.ID)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.booking.Update.Exec] Failed update Booking", log.Field("request", request), log.Field("error", err))
		return err
	}

	return nil
}
