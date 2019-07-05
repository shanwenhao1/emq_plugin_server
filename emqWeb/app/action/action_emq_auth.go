package action

import (
	"emq_plugin_server/emqWeb/app/service"
	"emq_plugin_server/emqWeb/domain/model"
	"github.com/gin-gonic/gin"
)

// Emq acl action struct
type EmqAclModel struct {
}

/* Emq web hook struct
可参考https://github.com/emqx/emqx-web-hook
*/
type EmqWebHookModel struct {
}

// emq auth 认证
func (this EmqAclModel) AuthAcl(c *gin.Context) {
	// 获取请求参数, 可以考虑在此添加中间件
	rjm := service.GetEmqReqData(c)
	if rjm != nil {
		jsonModel := rjm.(model.EmqAcl)
		// 可以做一些验证
		result := AuthAcl(jsonModel)
		service.EMQCommonResponse(c, result)
	}
}

// emq super acl 认证
func (this EmqAclModel) SuperAcl(c *gin.Context) {
	rjm := service.GetEmqReqData(c)
	if rjm != nil {
		jsonModel := rjm.(model.EmqAcl)
		// 可以做一些验证
		result := SuperAcl(jsonModel)
		service.EMQCommonResponse(c, result)
	}
}

// emq action acl 认证
func (this EmqAclModel) AclReq(c *gin.Context) {
	rjm := service.GetEmqReqData(c)
	if rjm != nil {
		jsonModel := rjm.(model.EmqAcl)
		// 可以做一些验证
		result := AclReq(jsonModel)
		service.EMQCommonResponse(c, result)
	}
}
