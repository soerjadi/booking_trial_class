package trial_class

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (r *trialClassRepository) Create(ctx context.Context, request domain.TrialClass) (trialClass domain.TrialClass, err error) {
	query := `
	INSERT INTO
		trial_classes (
			name,
			quota,
			available_slots
		)
	VALUES
		($1, $2, $3)
	RETURNING
		id,
		name,
		quota,
		available_slots
	`

	err = db.QuerierFromContext(ctx, r.db).QueryRow(ctx, query, request.Name, request.Quota, request.AvailableSlots).Scan(&trialClass.ID, &trialClass.Name, &trialClass.Quota, &trialClass.AvailableSlots)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.trial_class.Create.QueryRow] Failed create TrialClass", log.Field("request", request), log.Field("error", err))
		return domain.TrialClass{}, err
	}

	return
}
