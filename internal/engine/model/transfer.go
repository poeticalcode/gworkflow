package model

// TransferModel 转移节点模型，负责节点流转
type TransferModel struct {
	BaseModel
	// 要转移的节点
	Source *NodeModel
	// 转移的目标节点
	Target *NodeModel
	// 下一个节点 ID
	NextId string
	// 变迁的条件表达式
	Expr    string
	Enabled bool
}
