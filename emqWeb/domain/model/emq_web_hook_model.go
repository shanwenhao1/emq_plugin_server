package model

import (
	"emq_plugin_server/emqWeb/Infra/enum"
)

// web hook interface
type WebHook interface {
	ActionMethod() ParamModel
}

type WebHookAction struct {
	Action string `json:"action"` // action type.
}

type WebHookNormal struct {
	WebHookAction
	ClientId string `json:"client_id"` // client id
	Username string `json:"username"`  // username
}

type WHClientConn struct {
	WebHookNormal
	Ipaddress   string `json:"ipaddress"`    // client ipAddress
	Keepalive   int    `json:"keepalive"`    // long connect on http
	ProtoVer    int    `json:"proto_ver"`    // proto version
	ConnectedAt int    `json:"connected_at"` // connected timestamp

	ConnAck int `json:"conn_ack"` //
}

func (act WHClientConn) ActionMethod() ParamModel {
	// do something
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHClientDisConn struct {
	WebHookNormal
	Reason string `json:"reason"` // disconnected reason
}

func (act WHClientDisConn) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHClientSub struct {
	WebHookNormal
	Topic string                 `json:"topic"` // the topic to subscribe
	Opts  map[string]interface{} `json:"opts"`  // subscribe parameter
}

func (act WHClientSub) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHClientUnSub struct {
	WebHookNormal
	Topic string `json:"topic"` // the topic to unsubscribe
}

func (act WHClientUnSub) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHSessionCreate struct {
	WebHookNormal
}

func (act WHSessionCreate) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHSessionSub struct {
	WebHookNormal
	Topic string                 `json:"topic"` // the topic of session to subscribe
	Opts  map[string]interface{} `json:"opts"`  // subscribe parameter
}

func (act WHSessionSub) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHSessionUnSub struct {
	WebHookNormal
	Topic string `json:"topic"` // the topic of session to unsubscribe
}

func (act WHSessionUnSub) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHSessionT struct {
	WebHookNormal
	Reason string `json:"reason"` // session terminated reason
}

func (act WHSessionT) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHMsgPub struct {
	WebHookAction
	FromClientId string `json:"from_client_id"` // the clientId of message from
	FromUsername string `json:"from_username"`  // the client username of message from
	Topic        string `json:"topic"`          // the topic that message to publish
	Qos          int    `json:"qos"`            //
	Retain       bool   `json:"retain"`         // whether store the message or not (if it's true that the after subscriber can receive the message)
	Payload      string `json:"payload"`        // the message that you want to publish
	Ts           int    `json:"ts"`             // message publish timestamp
}

func (act WHMsgPub) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHMsgDeliver struct {
	WHMsgPub
}

func (act WHMsgDeliver) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

type WHMsgAck struct {
	WHMsgPub
}

func (act WHMsgAck) ActionMethod() ParamModel {
	return ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}
