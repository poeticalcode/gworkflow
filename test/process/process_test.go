package process

import (
	"testing"

	"github.com/poeticalcode/gworkflow/internal/engine"
)

// TestEngineDeploy 测试引擎发布
func TestProcessDeploy(t *testing.T) {
	wfe := engine.Default()
	wfe.Deploy(nil, "")
}
