package service

import (
	"errors"

	"github.com/beevik/etree"
	"github.com/poeticalcode/gworkflow/internal/engine/model"
)

// 流程相关服务
type ProcessService struct{}

// 通过 XML 文件配置一条流程
func (p ProcessService) DeployByXML(xmlData []byte, createBy string) error {
	if xmlData != nil {
		panic(errors.New("xml 数据不能为空"))
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xmlData); err != nil {
		panic(err)
	}
	p.parseModelXML(doc)
	return nil
}

// 将 XML 文件数据解析成流程模型
func (w ProcessService) parseModelXML(doc *etree.Document) (*model.ProcessModel, error) {
	// 获取流程节点
	root := doc.Root()
	process := &model.ProcessModel{}
	// 流程的 ID
	process.ID = root.SelectAttrValue(model.ATTR_ID, "")
	// 流程显示名称
	process.DisplayName = root.SelectAttrValue(model.ATTR_DISPLAYNAME, "")
	// 流程超时时间
	process.DisplayName = root.SelectAttrValue(model.ATTR_EXPIRETIME, "")
	// 流程节点下的所有子节点均需要转为节点模型
	for _, e := range root.ChildElements() {
		nodeModel := parseNodelModelXML(e)
		process.Nodes = append(process.Nodes, *nodeModel)
	}
	return nil, nil
}

// 将传入的 xml 元素转为节点模型
func parseNodelModelXML(e *etree.Element) *model.NodeModel {

	return nil
}
