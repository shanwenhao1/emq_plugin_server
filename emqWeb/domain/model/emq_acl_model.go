package model

type EmqAcl struct {
	ClientId string `json:"clientId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Topic    string `json:"topic"`
	Ip       string `json:"ip"`
	Access   string `json:"access"`
}
