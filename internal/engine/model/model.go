package model

const (
	ATTR_ID               = "id"
	ATTR_DISPLAYNAME      = "displayName"
	ATTR_INSTANCEURL      = "instanceUrl"
	ATTR_INSTANCENOCLASS  = "instanceNoClass"
	ATTR_EXPR             = "expr"
	ATTR_HANDLECLASS      = "handleClass"
	ATTR_FORM             = "form"
	ATTR_FIELD            = "field"
	ATTR_VALUE            = "value"
	ATTR_ATTR             = "attr"
	ATTR_TYPE             = "type"
	ATTR_ASSIGNEE         = "assignee"
	ATTR_ASSIGNEE_HANDLER = "assignmentHandler"
	ATTR_PERFORMTYPE      = "performType"
	ATTR_TASKTYPE         = "taskType"
	ATTR_TO               = "to"
	ATTR_PROCESSNAME      = "processName"
	ATTR_VERSION          = "version"
	ATTR_EXPIRETIME       = "expireTime"
	ATTR_AUTOEXECUTE      = "autoExecute"
	ATTR_CALLBACK         = "callback"
	ATTR_REMINDERTIME     = "reminderTime"
	ATTR_REMINDERREPEAT   = "reminderRepeat"
	ATTR_CLAZZ            = "clazz"
	ATTR_METHODNAME       = "methodName"
	ATTR_ARGS             = "args"
	ATTR_VAR              = "var"
	ATTR_LAYOUT           = "layout"
	ATTR_G                = "g"
	ATTR_OFFSET           = "offset"
	ATTR_PREINTERCEPTORS  = "preInterceptors"
	ATTR_POSTINTERCEPTORS = "postInterceptors"
)

// 基础模型，所有的模型都要嵌入它
type BaseModel struct {
	ID          string
	DisplayName string
}

// 节点模型
type NodeModel struct {
	BaseModel
}
