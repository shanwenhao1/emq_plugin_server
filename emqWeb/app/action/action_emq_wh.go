package action

import (
	"emq_plugin_server/emqWeb/Infra/enum"
	"emq_plugin_server/emqWeb/app/service"
	"emq_plugin_server/emqWeb/domain/model"
	"github.com/gin-gonic/gin"
)

// emq web hook action
func (this EmqWebHookModel) WebHook(c *gin.Context) {
	/*	获取请求json数据
		先根据json action type判别request 请求类型
	*/
	rjm := service.GetEmqWebHookData(c)
	if rjm != nil {
		result := WebHookFunc(rjm)
		service.EMQCommonResponse(c, result)
	} else {
		result := model.ParamModel{ErrorCode: enum.OPERATE_FAILED}
		service.EMQCommonResponse(c, result)
	}
}
