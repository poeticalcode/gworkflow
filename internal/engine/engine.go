package engine

import "github.com/poeticalcode/gworkflow/internal/engine/service"

// 定义一个引擎
type WorkFlowEngine struct {
	ProcessService service.ProcessService
}

// 实例化一个默认的工作流引擎
func Default() *WorkFlowEngine {
	return &WorkFlowEngine{}
}
