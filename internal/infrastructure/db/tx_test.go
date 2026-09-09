package db

import (
	"context"
	"errors"
	"testing"

	"github.com/soerjadi/booking/internal/infrastructure/config"
)

// testDSN matches config.toml's [database] section (local dev Postgres).
const testDSN = "postgres://postgres:postgres@localhost:5432/cbt?sslmode=disable"

func mustTestPool(t *testing.T) {
	t.Helper()
	if pool != nil {
		return
	}
	if _, err := NewPGXPool(context.Background(), config.DBConfig{
		URL:             testDSN,
		MaxConns:        5,
		MinConns:        1,
		MaxConnLifetime: 30,
		MaxConnIdleTime: 5,
	}); err != nil {
		t.Skipf("no local postgres available at %s: %v", testDSN, err)
	}
}

// TestDo_RollsBackWritesMadeThroughQuerierFromContext is the regression
// guard for the Do()/WithTx/QuerierFromContext contract: a write issued via
// QuerierFromContext(ctx, fallback) inside a Do() block must be routed
// through the tx Do() opened, so it's undone when fn returns an error.
func TestDo_RollsBackWritesMadeThroughQuerierFromContext(t *testing.T) {
	mustTestPool(t)
	ctx := context.Background()

	if _, err := pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS tx_test_scratch (id serial primary key, note text)`); err != nil {
		t.Fatalf("create scratch table: %v", err)
	}
	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DROP TABLE IF EXISTS tx_test_scratch`)
	})

	sentinel := errors.New("forced failure after write")
	err := Do(ctx, func(ctx context.Context) error {
		// Mirrors what a repository does: fetch the active executor from
		// ctx instead of writing via the pool directly.
		if _, err := QuerierFromContext(ctx, pool).Exec(ctx, `INSERT INTO tx_test_scratch (note) VALUES ('should be rolled back')`); err != nil {
			return err
		}
		return sentinel
	})
	if !errors.Is(err, sentinel) {
		t.Fatalf("expected sentinel error, got %v", err)
	}

	var count int
	if err := pool.QueryRow(ctx, `SELECT count(*) FROM tx_test_scratch`).Scan(&count); err != nil {
		t.Fatalf("count scratch rows: %v", err)
	}

	if count != 0 {
		t.Fatalf("Do() reported an error but the write survived: found %d row(s) in tx_test_scratch; "+
			"a write made via QuerierFromContext(ctx, fallback) should be routed through Do()'s tx and rolled back with it", count)
	}
}
