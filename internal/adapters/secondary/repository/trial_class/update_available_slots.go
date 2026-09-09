package trial_class

import (
	"context"
	"errors"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

var ErrNoAvailableSlots = errors.New("no available slots")

func (r *trialClassRepository) UpdateAvailableSlots(ctx context.Context, request domain.TrialClass) error {
	query := `
	UPDATE
		trial_classes
	SET
		available_slots = $1
	WHERE
		id = $2
		AND available_slots > 0
	`

	tag, err := db.QuerierFromContext(ctx, r.db).Exec(ctx, query, request.AvailableSlots, request.ID)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.trial_class.UpdateAvailableSlots.Exec] Failed update TrialClass available slots", log.Field("request", request), log.Field("error", err))
		return err
	}

	if tag.RowsAffected() == 0 {
		log.ErrorCtx(ctx, "[repository.trial_class.UpdateAvailableSlots.Exec] No rows affected, class not found or no available slots", log.Field("request", request))
		return ErrNoAvailableSlots
	}

	return nil
}
