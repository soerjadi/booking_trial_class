package trial_class

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (r *trialClassRepository) GetMember(ctx context.Context, idClass int64) ([]domain.TrialClassMember, error) {
	var members []domain.TrialClassMember
	query := `
		SELECT id, trial_classes_id, student_id 
		FROM trial_class_members 
		WHERE trial_classes_id = $1
	`
	rows, err := db.QuerierFromContext(ctx, r.db).Query(ctx, query, idClass)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.trial_class.GetMember.Query] Failed get TrialClassMember", log.Field("idClass", idClass), log.Field("error", err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member domain.TrialClassMember
		if err := rows.Scan(&member.ID, &member.TrialClassID, &member.StudentID); err != nil {
			log.ErrorCtx(ctx, "[repository.trial_class.GetMember.Scan] Failed scan TrialClassMember", log.Field("error", err))
			return nil, err
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		log.ErrorCtx(ctx, "[repository.trial_class.GetMember.RowsErr] Error iterating TrialClassMember rows", log.Field("error", err))
		return nil, err
	}

	return members, nil
}
