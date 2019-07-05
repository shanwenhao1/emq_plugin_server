package service

import (
	"emq_plugin_server/emqWeb/Infra/enum"
	"emq_plugin_server/emqWeb/domain/model"
	"encoding/json"
	"errors"
)

// 根据action choose返回响应的struct结构体
func GetWHActData(data *[]byte, actChoose string) (interface{}, error) {
	var actionType int
	var err error
	// 未找到action type
	if value, ok := enum.WHActMap[actChoose]; !ok {
		return nil, errors.New(enum.EMQStatMap[enum.ACTION_DEFECT])
	} else {
		actionType = value
	}
	// 返回指定类型
	switch actionType {
	case enum.WHCliConn:
		var actData model.WHClientConn
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHCliDisConn:
		var actData model.WHClientDisConn
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHCliSub:
		var actData model.WHClientSub
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHCliUnSub:
		var actData model.WHClientUnSub
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHSesCre:
		var actData model.WHSessionCreate
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHSesSub:
		var actData model.WHSessionSub
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHSesUnSub:
		var actData model.WHSessionUnSub
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHSesT:
		var actData model.WHSessionT
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHMsgPub:
		var actData model.WHMsgPub
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHMsgDeliver:
		var actData model.WHMsgDeliver
		err = json.Unmarshal(*data, &actData)
		return actData, err
	case enum.WHMsgAck:
		var actData model.WHMsgAck
		err = json.Unmarshal(*data, &actData)
		return actData, err
	default:
		return nil, errors.New(enum.EMQStatMap[enum.ACTION_FAILED])
	}
}
