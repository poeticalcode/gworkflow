package process

import (
	"io"
	"os"
	"testing"

	"github.com/poeticalcode/gworkflow/internal/engine"
)

// TestProcessDeploy 测试引擎发布
func TestProcessDeploy(t *testing.T) {
	e := engine.Default()
	ps := e.ProcessService

	file, err := os.Open("./process.xml")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	xmlData, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}
	ps.DeployByXML(xmlData, "he.wenyao")
}
