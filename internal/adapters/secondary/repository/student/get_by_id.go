package student

import (
	"context"

	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/domain"
)

func (r *studentRepository) GetByID(ctx context.Context, id int64) (student domain.Student, err error) {
	query := `
	SELECT
		id,
		name,
		parent_id
	FROM 
		students
	WHERE 
		id = $1
	LIMIT 1
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&student.ID, &student.Name, &student.ParentID)
	if err != nil {
		log.ErrorCtx(ctx, "[repository.student.GetByID] failed get student by id", log.Field("request", id), log.Field("error", err))
		return domain.Student{}, err
	}
	return
}
