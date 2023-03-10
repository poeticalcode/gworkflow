package process_test

import (
	"io"
	"os"
	"testing"

	"github.com/poeticalcode/gworkflow/internal/engine"
)

// TestProcessDeploy 测试引擎发布
func TestProcessDeploy(t *testing.T) {
	//e := engine.Default()
	//ps := e.ProcessService
	file, err := os.Open("./process.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	xmlData, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(xmlData))

	e := engine.Default()

	err = e.ProcessService.DeployByXML(xmlData, "he.wenyao")
	if err != nil {
		t.Fatal(err)
	}
}
