package parser

import (
	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
)

// EndParser 结束节点解析器
type EndParser struct {
	BaseParser
}

// parse 将 xml 节点转为基础的节点模型
func (end EndParser) parse(e etree.Document) *model.NodeModel {
	// 初始化一个
	end.model = &model.NodeModel{}
	// 调用公共方法
	end.BaseParser.parse(e)
	return end.model
}
