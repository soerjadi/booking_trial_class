package parent

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *parentRepository) GetByID(ctx context.Context, id int64) (parent domain.Parent, err error) {
	query := `
	SELECT
		id,
		name
	FROM 
		parents
	WHERE 
		id = $1
	LIMIT 1
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&parent.ID, &parent.Name)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.parent.GetByID] failed get parent by id", log.Field("request", id), log.Field("error", err))
		return domain.Parent{}, err
	}
	return
}
