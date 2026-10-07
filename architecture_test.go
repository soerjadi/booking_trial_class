package booking_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArchitectureCopyright(t *testing.T) {
	data, err := os.ReadFile("architecture.md")
	require.NoError(t, err)
	content := string(data)
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")

	assert.GreaterOrEqual(t, len(lines), 137)
	assert.Contains(t, content, "## Copyright")
	assert.Contains(t, content, "© 2026 All rights reserved.")

	idx := strings.Index(content, "Last updated:")
	require.NotEqual(t, -1, idx)
	ts := strings.TrimSpace(content[idx+len("Last updated:"):])
	ts = strings.Fields(ts)[0]
	_, err = time.Parse(time.RFC3339Nano, ts)
	if err != nil {
		_, err = time.Parse(time.RFC3339, ts)
	}
	require.NoError(t, err)

	assert.Contains(t, content, "# Architecture Decision")
	assert.Contains(t, content, "## Flow Diagram")
}
