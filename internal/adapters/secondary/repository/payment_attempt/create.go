package payment_attempt

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *paymentAttemptRepository) Create(ctx context.Context, request domain.PaymentAttempt) (paymentAttempt domain.PaymentAttempt, err error) {
	query := `
	INSERT INTO
		payment_attempts (
			booking_id,
			status,
			note,
			created_at
		)
	VALUES
		($1, $2, $3, NOW())
	RETURNING
		id,
		booking_id,
		status,
		note,
		created_at,
		updated_at
	`

	err = r.db.QueryRow(ctx, query, request.BookingID, request.Status, request.Note).Scan(
		&paymentAttempt.ID, &paymentAttempt.BookingID, &paymentAttempt.Status, &paymentAttempt.Note, &paymentAttempt.CreatedAt, &paymentAttempt.UpdatedAt,
	)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.payment_attempt.Create.QueryRow] Failed create PaymentAttempt", log.Field("request", request), log.Field("error", err))
		return domain.PaymentAttempt{}, err
	}

	return
}
