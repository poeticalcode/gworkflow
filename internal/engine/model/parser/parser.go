package parser

import (
	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
)

// NodeParser 解析节点
type NodeParser interface {
	parse(e etree.Document) *model.NodeModel
}

// BaseParser 基础的解析器
type BaseParser struct {
	model *model.NodeModel
}

// parse 将 xml 节点转为基础的节点模型
func (b BaseParser) parse(e etree.Document) *model.NodeModel {

	// 自定义解析
	return b.model
}
