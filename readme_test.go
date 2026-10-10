package booking_test

import (
	"os"
	"strings"
	"testing"
)

func TestReadmeCopyrightNotice(t *testing.T) {
	b, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("read README.md: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, "*This README was updated by an AI agent on 2026-10-10T01:56:00.000Z.*") {
		t.Fatalf("copyright notice not found")
	}
	if !strings.HasSuffix(strings.TrimSpace(s), "*This README was updated by an AI agent on 2026-10-10T01:56:00.000Z.*") {
		t.Fatalf("copyright notice not at end of file")
	}
	if !strings.Contains(s, "---\n\n*This README") {
		t.Fatalf("expected separator before copyright notice")
	}
}
