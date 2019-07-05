package service

import (
	"emq_plugin_server/emqWeb/Infra/log"
	"emq_plugin_server/emqWeb/domain/model"
	"emq_plugin_server/emqWeb/domain/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"strings"
)

/*
	Emq web hook request data analysis
*/
func GetEmqWebHookData(c *gin.Context) interface{} {
	req := c.Request
	addr := req.Header.Get("X-Real-IP") // 获取真实发出请求的客户端IP
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For") // 获取IP(包含代理IP）
		if addr == "" {
			addr = req.RemoteAddr
		}
	}
	//log.LogWithTag(log.InfoLog, log.ReqParse, "Emq Request %s for %s", req.URL.Path, addr)
	dataS, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.LogWithTag(log.ErrorLog, log.ReqParse, "%w : %w", "Gin Read Body Error", err)
		return nil
	}
	log.LogWithTag(log.InfoLog, log.ReqParse, "%v : %v", "The Emq Web Hook Request Body is", string(dataS))

	var act model.WebHookAction
	err = json.Unmarshal(dataS, &act)
	if err != nil {
		log.LogWithTag(log.ErrorLog, log.ReqParse, "%v : %v", "Convert Body To Json Failed", err)
		return nil
	}
	// 调用web hook领域内的服务以解析数据
	actionSTRT, err := service.GetWHActData(&dataS, act.Action)
	if err != nil {
		log.LogWithTag(log.ErrorLog, log.ReqParse, "Action choose Failed, err: %w", err)
		return nil
	}
	return actionSTRT
}

/*
	Emq http auth acl request data analysis
*/
func GetEmqReqData(c *gin.Context) interface{} {
	var reqData model.EmqAcl
	req := c.Request
	addr := req.Header.Get("X-Real-IP") // 获取真实发出请求的客户端IP
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For") // 获取IP(包含代理IP）
		if addr == "" {
			addr = req.RemoteAddr
		}
	}
	//log.LogWithTag(log.InfoLog, log.ReqParse, "Emq Request %s for %s", req.URL.Path, addr)
	dataS, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.LogWithTag(log.ErrorLog, log.ReqParse, "%w : %w", "Gin Read Body Error", err)
	}
	log.LogWithTag(log.InfoLog, log.ReqParse, "%v : %v", "The Emq Request Body is", string(dataS))

	// 将http参数(clientid=c5c315ec-ecba-4d5a-b1ae-7adaf6e1565b1561965833405&username=admin)转换成map[string]interface形式
	reqMap := StrPar(string(dataS))
	// 将map装换为指定的结构体
	if err := mapstructure.Decode(reqMap, &reqData); err != nil {
		log.LogWithTag(log.ErrorLog, log.ReqParse, "Convert request map to struct failed, map: %s, err: %w",
			string(dataS), err)
		EmqResponseData(c, false)
		return nil
	}
	return reqData
}

/* 获取emq http request参数,
reqStr格式: clientid=c5c315ec-ecba-4d5a-b1ae-7adaf6e1565b1561973576671&username=admin&password=public
return: map[clientid:c5c315ec-ecba-4d5a-b1ae-7adaf6e1565b1561973576671 username:admin password:public]
*/
func StrPar(reqStr string) map[string]interface{} {
	splitStr := strings.SplitAfter(reqStr, "&")
	parMap := make(map[string]interface{}, len(splitStr))
	for _, meteData := range splitStr {
		newMeteData := strings.Replace(meteData, "&", "", 1)
		splitMD := strings.Split(newMeteData, "=")
		parMap[splitMD[0]] = splitMD[1]
	}
	//strHex, err := regexp.Compile(`((.*?=.*?[?=&])|([?=&].*?=.*?))`)
	//if err != nil{
	//	panic(err)
	//}
	//result2 := strHex.FindAllString(reqStr, -1)
	return parMap
}
