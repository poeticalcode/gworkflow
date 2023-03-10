package parser

import (
	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
)

// StartParser 开始节点解析器
type StartParser struct {
	baseParser
}

// 校验 EndParser 是否实现了 NodeParser 接口
var _ NodeParser = (*StartParser)(nil)

// parse 将 xml 节点转为基础的节点模型
func (s StartParser) Parse(e *etree.Element) *model.NodeModel {
	// 初始化一个
	s.model = &model.NodeModel{}
	// 调用公共方法
	s.baseParser.Parse(e)
	return s.model
}
