package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	log "github.com/oemahdev/logger"
	"github.com/soerjadi/booking/internal/core/ports/repository"
)

type txCtxKey struct{}

// WithTx returns a context carrying q as the active transaction executor.
func WithTx(ctx context.Context, q repository.DB) context.Context {
	return context.WithValue(ctx, txCtxKey{}, q)
}

// QuerierFromContext returns the transaction executor stored in ctx by
// WithTx, or fallback if ctx carries none. Repositories should call this
// instead of using their injected DB directly, so writes issued inside a
// Do block participate in that transaction.
func QuerierFromContext(ctx context.Context, fallback repository.DB) repository.DB {
	if q, ok := ctx.Value(txCtxKey{}).(repository.DB); ok {
		return q
	}
	return fallback
}

// Do runs fn inside a database transaction on the pool created by
// NewPGXPool, committing if fn returns nil and rolling back otherwise
// (including on panic). Call it directly from services; no injected
// dependency is needed.
var Do = func(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.ErrorCtx(ctx, "[infrastructure.db.BeginTx] failed to start", log.Field("error", err))
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			log.ErrorCtx(ctx, "[infrastructure.db.Rollback] rollback transaction", log.Field("error", err), log.Field("panic", p))
		} else if err != nil {
			_ = tx.Rollback(ctx)
			log.ErrorCtx(ctx, "[infrastructure.db.Rollback] rollback transaction", log.Field("error", err), log.Field("panic", p))
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = fn(WithTx(ctx, tx))
	return err
}
