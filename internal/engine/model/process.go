package model

// 流程模型
type ProcessModel struct {
	BaseModel
	// 流程中所有的节点
	Nodes []NodeModel
}
