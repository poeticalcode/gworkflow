package engine

import (
	"io"
)

// 定义一个引擎
type WorkFlowEngine struct {
}

// 实例化一个默认的工作流引擎
func Default() *WorkFlowEngine {
	return &WorkFlowEngine{}
}

// 配置一条工作流
func (w WorkFlowEngine) Deploy(r io.Reader, createBy string) error {
	// if r != nil {
	// 	panic(errors.New())
	// }
	return nil
}
