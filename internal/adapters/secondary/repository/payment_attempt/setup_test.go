package payment_attempt

import (
	"testing"
	"time"

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
		case *domain.PaymentAttemptStatus:
			*d = v.(domain.PaymentAttemptStatus)
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
