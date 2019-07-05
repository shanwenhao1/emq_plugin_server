package enum

const (
	WHCliConn    = iota // client connect action
	WHCliDisConn        // client disconnect action
	WHCliSub            // client subscribe action
	WHCliUnSub          // client unsubscribe action
	WHSesCre            // session create action
	WHSesSub            // session subscribe action
	WHSesUnSub          // session unsubscribe action
	WHSesT              // session terminated action
	WHMsgPub            // message publish action
	WHMsgDeliver        // message deliver action
	WHMsgAck            // message ack action
)

var WHActMap map[string]int = map[string]int{
	"client_connected":     WHCliConn,
	"client_disconnected":  WHCliDisConn,
	"client_subscribe":     WHCliSub,
	"client_unsubscribe":   WHCliUnSub,
	"session_created":      WHSesCre,
	"session_subscribed":   WHSesSub,
	"session_unsubscribed": WHSesUnSub,
	"session_terminated":   WHSesT,
	"message_publish":      WHMsgPub,
	"message_delivered":    WHMsgDeliver,
	"message_acked":        WHMsgAck,
}

const (
	ACTION_SUCCESS = iota
	ACTION_FAILED
	ACTION_DEFECT
)

var EMQStatMap = map[int]string{
	ACTION_DEFECT: "action类型错误",
	ACTION_FAILED: "action操作失败",
}
