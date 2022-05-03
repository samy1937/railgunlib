package exp_templates

/*
ExmTemplate是所有exp的模板
编写实际exp时，只需内嵌这个结构体即可，各功能exp还是按以往规律，签名如下
func (self *exp) GetMsg1(cmd string) (expResult templates.ExpResult)
func (self *exp) Cmd1(cmd string) (expResult templates.ExpResult)
func (self *exp) Reverse1(ip, port string) (expResult templates.ExpResult)
func (self *exp) Upload1(filename string, content string) (status bool, respPath string)

Init方法用于初始化
EchoMsgWithReturn 用于实时回显，判断TmpResultCh是否不为空，否则即打印log

*/

import (
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
)

type ExpTemplate struct {
	Params exp_model.ExpSendParamS
	//Num2AliasMap map[string] string
}

// 获取设置的http头部
func (self *ExpTemplate) GetInitExpHeaders() (headers lzhttp.Header) {
	return
}

// 安全获取map里的value
func (self *ExpTemplate) GetItemSafe(s map[string]string, key string) (ret string) {
	return
}

// 安全添加URI
func (self *ExpTemplate) AddUri(target, uri string) (result string) {

	return
}

// 追加URL，基于当前目录
func (self *ExpTemplate) AppendUri(target, uri string) (result string) {
	return
}
func (self *ExpTemplate) GetHostname(target string) (hostname string) {
	return
}
