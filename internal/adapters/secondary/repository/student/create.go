package student

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
	"github.com/soerjadi/booking/internal/infrastructure/db"
)

func (r *studentRepository) Create(ctx context.Context, request domain.Student) (student domain.Student, err error) {
	query := `
	INSERT INTO
		students (
			name,
			parent_id
		)
	VALUES
		($1, $2)
	RETURNING
		id,
		name,
		parent_id
	`

	err = db.QuerierFromContext(ctx, r.db).QueryRow(ctx, query, request.Name, request.ParentID).Scan(&student.ID, &student.Name, &student.ParentID)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.student.Create.QueryRow] Failed create Students", log.Field("request", request), log.Field("error", err))
		return domain.Student{}, err
	}

	return
}
