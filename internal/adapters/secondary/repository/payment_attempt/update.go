package payment_attempt

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *paymentAttemptRepository) Update(ctx context.Context, request domain.PaymentAttempt) error {
	query := `
	UPDATE
		payment_attempts
	SET
		booking_id = $1,
		status = $2,
		note = $3,
		created_at = $4,
		updated_at = $5
	WHERE
		id = $6
	`

	_, err := r.db.Exec(ctx, query, request.BookingID, request.Status, request.Note, request.CreatedAt, request.UpdatedAt, request.ID)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.payment_attempt.Update.Exec] Failed update PaymentAttempt", log.Field("request", request), log.Field("error", err))
		return err
	}

	return nil
}
