package trial_class

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *trialClassRepository) Update(ctx context.Context, request domain.TrialClass) error {
	query := `
	UPDATE
		trial_classes
	SET
		name = $1,
		quota = $2,
		available_slots = $3
	WHERE
		id = $4
	`

	_, err := r.db.Exec(ctx, query, request.Name, request.Quota, request.AvailableSlots, request.ID)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.trial_class.Update.Exec] Failed update TrialClass", log.Field("request", request), log.Field("error", err))
		return err
	}

	return nil
}
