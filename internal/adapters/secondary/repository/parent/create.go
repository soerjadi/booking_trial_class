package parent

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (r *parentRepository) Create(ctx context.Context, request domain.Parent) (parent domain.Parent, err error) {
	query := `
	INSERT INTO
		parents (
			name
		)
	VALUES
		($1)
	RETURNING
		id,
		name
	`

	err = db.QuerierFromContext(ctx, r.db).QueryRow(ctx, query, request.Name).Scan(&parent.ID, &parent.Name)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.parent.Create.QueryRow] Failed create Parent", log.Field("request", request), log.Field("error", err))
		return domain.Parent{}, err
	}

	return
}
