package model

type CreateRedis struct {
	ServiceName string `json:"serviceName"`
	Prefix      string `json:"prefix"`
	JsonName    string `json:"jsonName"`
}
