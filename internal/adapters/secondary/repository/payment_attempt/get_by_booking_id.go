package payment_attempt

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (r *paymentAttemptRepository) GetByBookingID(ctx context.Context, bookingID int64) (domain.PaymentAttempt, error) {
	var pa domain.PaymentAttempt
	query := `
		SELECT id, booking_id, status, note, created_at, updated_at 
		FROM payment_attempts 
		WHERE booking_id = $1
	`
	err := db.QuerierFromContext(ctx, r.db).QueryRow(ctx, query, bookingID).Scan(
		&pa.ID, &pa.BookingID, &pa.Status, &pa.Note, &pa.CreatedAt, &pa.UpdatedAt,
	)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.paymentattempt.GetByBookingID.QueryRow] Failed get payment attempt by booking id", log.Field("bookingID", bookingID), log.Field("error", err))
		return domain.PaymentAttempt{}, err
	}
	return pa, nil
}
