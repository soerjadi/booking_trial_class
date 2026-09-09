package booking

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (r *bookingRepository) GetByPaymentCode(ctx context.Context, paymentCode string) (domain.Booking, error) {
	var b domain.Booking
	query := `
		SELECT id, trial_classes_id, student_id, status, idempotency_key, hold_expires_at, payment_code, created_at, updated_at 
		FROM bookings 
		WHERE payment_code = $1
	`
	err := db.QuerierFromContext(ctx, r.db).QueryRow(ctx, query, paymentCode).Scan(
		&b.ID,
		&b.TrialClassID,
		&b.StudentID,
		&b.Status,
		&b.IdempotencyKey,
		&b.HoldExpiredAt,
		&b.PaymentCode,
		&b.CreatedAt,
		&b.UpdatedAt,
	)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.booking.GetByPaymentCode.QueryRow] Failed get booking by payment code", log.Field("paymentCode", paymentCode), log.Field("error", err))
		return domain.Booking{}, err
	}

	return b, nil
}
