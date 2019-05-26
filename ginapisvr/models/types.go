package models

type ResHeader struct {
	Code uint32 `form:"code" json:"code"`
	Msg string `form:"msg" json:"msg"`
}