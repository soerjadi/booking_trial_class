package booking

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (r *bookingRepository) Create(ctx context.Context, request domain.Booking) (booking domain.Booking, err error) {
	query := `
	INSERT INTO
		bookings (
			trial_classes_id,
			student_id,
			status,
			idempotency_key,
			hold_expires_at,
			payment_code,
			created_at
		)
	VALUES
		($1, $2, $3, $4, $5, $6, NOW())
	RETURNING
		id,
		trial_classes_id,
		student_id,
		status,
		idempotency_key,
		hold_expires_at,
		payment_code,
		created_at,
		updated_at
	`

	err = db.QuerierFromContext(ctx, r.db).QueryRow(ctx, query, request.TrialClassID, request.StudentID, request.Status, request.IdempotencyKey, request.HoldExpiredAt, request.PaymentCode).Scan(
		&booking.ID, &booking.TrialClassID, &booking.StudentID, &booking.Status, &booking.IdempotencyKey, &booking.HoldExpiredAt, &booking.PaymentCode, &booking.CreatedAt, &booking.UpdatedAt,
	)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.booking.Create.QueryRow] Failed create Booking", log.Field("request", request), log.Field("error", err))
		return domain.Booking{}, err
	}

	return
}
