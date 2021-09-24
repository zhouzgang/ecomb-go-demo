package model

type Result struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
