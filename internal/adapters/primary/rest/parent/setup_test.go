package parent

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func setupMocks(t *testing.T) *gomock.Controller {
	return gomock.NewController(t)
}
