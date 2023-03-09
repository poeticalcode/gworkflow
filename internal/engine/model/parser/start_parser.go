package parser

import (
	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
)

// StartParser 开始节点解析器
type StartParser struct {
	BaseParser
}

// parse 将 xml 节点转为基础的节点模型
func (s StartParser) parse(e etree.Document) *model.NodeModel {
	// 初始化一个
	s.model = &model.NodeModel{}
	// 调用公共方法
	s.BaseParser.parse(e)
	return s.model
}
