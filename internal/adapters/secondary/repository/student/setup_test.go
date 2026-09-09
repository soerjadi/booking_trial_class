package student

import (
	"time"
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
		case *time.Time:
			*d = v.(time.Time)
		}
	}
	return nil
}
