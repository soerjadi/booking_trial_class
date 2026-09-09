package trial_class

import (
	"context"

	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *trialClassRepository) GetByID(ctx context.Context, id int64) (domain.TrialClass, error) {
	var trialClass domain.TrialClass
	query := `
		SELECT id, name, quota, available_slots 
		FROM trial_classes 
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&trialClass.ID,
		&trialClass.Name,
		&trialClass.Quota,
		&trialClass.AvailableSlots,
	)
	if err != nil {
		return domain.TrialClass{}, err
	}

	return trialClass, nil
}
