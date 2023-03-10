package service

import (
	"errors"
	"fmt"

	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
	"github.com/poeticalcode/gworkflow/internal/engine/model/parser"
)

// ProcessService 流程相关服务
type ProcessService struct{}

// DeployByXML 通过 XML 文件配置一条流程
func (w ProcessService) DeployByXML(xmlData []byte, createBy string) error {
	if xmlData == nil {
		panic(errors.New("xml 数据不能为空"))
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xmlData); err != nil {
		panic(err)
	}
	_, err := w.parseModelXML(doc)
	if err != nil {
		return err
	}
	return nil
}

// 将 XML 文件数据解析成流程模型
func (w ProcessService) parseModelXML(doc *etree.Document) (*model.ProcessModel, error) {
	// 获取流程节点
	root := doc.Root()
	process := &model.ProcessModel{
		Nodes: []*model.NodeModel{},
	}
	// 流程的 ID
	process.ID = root.SelectAttrValue(model.ATTR_ID, "")
	// 流程显示名称
	process.DisplayName = root.SelectAttrValue(model.ATTR_DISPLAYNAME, "")
	// 流程超时时间
	process.DisplayName = root.SelectAttrValue(model.ATTR_EXPIRETIME, "")
	// 流程节点下的所有子节点均需要转为节点模型
	for _, e := range root.ChildElements() {
		nodeModel, _ := parseNodelModelXML(e)
		process.Nodes = append(process.Nodes, nodeModel)
	}

	// 遍历所有节点的输出节点，补全 inputs , target
	for _, node := range process.Nodes {
		for _, out := range node.Outputs {
			for _, node2 := range process.Nodes {
				if node2.ID == out.ID {
					node2.Inputs = append(node2.Inputs, out)
					out.Target = node2
				}
			}
		}
	}
	fmt.Println(process)
	return nil, nil
}

// 将传入的 xml 元素转为节点模型
func parseNodelModelXML(e *etree.Element) (*model.NodeModel, error) {
	// 根据节点名称获取节点解析器
	var p, err = parser.GetParser(e.Tag)
	if err != nil {
		return nil, err
	}
	return p.Parse(e), nil
}
