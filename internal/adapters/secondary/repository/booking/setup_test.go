package booking

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/soerjadi/booking/internal/core/domain"
	"go.uber.org/mock/gomock"
)

type mockRow struct {
	err  error
	args []any
}

func (m mockRow) Scan(dest ...any) error {
	if m.err != nil {
		return m.err
	}
	for i, v := range m.args {
		switch d := dest[i].(type) {
		case *int64:
			*d = v.(int64)
		case *string:
			*d = v.(string)
		case *domain.BookingStatus:
			*d = v.(domain.BookingStatus)
		case *time.Time:
			if v != nil {
				*d = v.(time.Time)
			}
		}
	}
	return nil
}

func setupMocks(t *testing.T) (*gomock.Controller, *mockRow) {
	ctrl := gomock.NewController(t)
	return ctrl, &mockRow{}
}

type mockRows struct {
	rows [][]any
	err  error
	pos  int
}

func (m *mockRows) Close()                                       {}
func (m *mockRows) Err() error                                   { return m.err }
func (m *mockRows) CommandTag() pgconn.CommandTag                { return pgconn.CommandTag{} }
func (m *mockRows) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (m *mockRows) Next() bool {
	if m.pos < len(m.rows) {
		m.pos++
		return true
	}
	return false
}
func (m *mockRows) Scan(dest ...any) error {
	if m.err != nil {
		return m.err
	}
	row := m.rows[m.pos-1]
	for i, v := range row {
		switch d := dest[i].(type) {
		case *int64:
			*d = v.(int64)
		case *string:
			*d = v.(string)
		case *domain.BookingStatus:
			*d = v.(domain.BookingStatus)
		case **time.Time:
			if v != nil {
				*d = v.(*time.Time)
			}
		case *time.Time:
			if v != nil {
				*d = v.(time.Time)
			}
		}
	}
	return nil
}
func (m *mockRows) Values() ([]any, error) { return nil, nil }
func (m *mockRows) RawValues() [][]byte    { return nil }
func (m *mockRows) Conn() *pgx.Conn        { return nil }
func (m *mockRows) TypeMap() *pgtype.Map   { return nil }
