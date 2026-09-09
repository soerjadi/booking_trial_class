package trial_class

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *trialClassRepository) InsertMember(ctx context.Context, request domain.TrialClassMember) error {
	query := `
	INSERT INTO
		trial_class_members (
			trial_classes_id,
			student_id
		)
	VALUES
		($1, $2)
	`

	_, err := r.db.Exec(ctx, query, request.TrialClassID, request.StudentID)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.trial_class.InsertMember.Exec] Failed insert TrialClassMember", log.Field("request", request), log.Field("error", err))
		return err
	}

	return nil
}
