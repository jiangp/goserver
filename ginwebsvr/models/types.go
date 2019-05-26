package models

type ResHeader struct {
	Code uint32 `form:"code" json:"code"`
	Msg string `form:"msg" json:"msg"`
}

//VersionInfo
type VersionInfo struct {
	CurVersion uint32
	VarTitle string
	VarMessage string
	IsForce uint32
	UpdateUrl string
}

type ResNotification struct {
	Id uint32
	Title string
	Content string
	Starttime int64
	Endtime int64
}