package parser

import (
	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
)

// EndParser 结束节点解析器
type EndParser struct {
	baseParser
}

// 校验 EndParser 是否实现了 NodeParser 接口
var _ NodeParser = (*EndParser)(nil)

// parse 将 xml 节点转为基础的节点模型
func (end EndParser) Parse(e *etree.Element) *model.NodeModel {
	// 初始化一个
	end.model = &model.NodeModel{}
	// 调用公共方法
	end.baseParser.Parse(e)
	return end.model
}
