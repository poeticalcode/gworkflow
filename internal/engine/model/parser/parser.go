package parser

import (
	"fmt"
	"sync"

	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
)

// NodeParser 解析节点
type NodeParser interface {
	Parse(e *etree.Element) *model.NodeModel
}

// baseParser 基础的解析器
type baseParser struct {
	model *model.NodeModel
}

// parse 将 xml 节点转为基础的节点模型
func (b baseParser) Parse(e *etree.Element) *model.NodeModel {
	b.model.DisplayName = e.SelectAttrValue(model.ATTR_DISPLAYNAME, "")
	b.model.ID = e.SelectAttrValue(model.ATTR_ID, "")
	// 自定义解析
	return b.model
}

var parsers sync.Map

// 获取解析器
func GetParser(nodeName string) (NodeParser, error) {
	parserIf, ok := parsers.Load(nodeName)
	if ok {
		return parserIf.(NodeParser), nil
	}
	var parser NodeParser
	switch nodeName {
	case "start":
		parser = startParser{}
	case "end":
		parser = endParser{}
	default:
		return nil, fmt.Errorf("unsupported node type %s", nodeName)
	}
	parsers.Store(nodeName, parser)
	return parser, nil
}
