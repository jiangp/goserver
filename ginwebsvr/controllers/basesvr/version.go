package basesvr

import (
	"encoding/json"
	"ginwebsvr/ginit"
	"ginwebsvr/models"
	"net/http"
	"szprotobuf"

	"github.com/gin-gonic/gin"
)

func CheckVsion(c *gin.Context) {

	var Reqdata szprotobuf.ReqVersion
	if err := c.ShouldBind(&Reqdata); err == nil{
		ginit.Scrlog.Info("VisitLogin Logintype:%v Type:%v, Version:%v", Reqdata.Logintype, Reqdata.Plattype, Reqdata.Version)
		Resdata := HandleFunc(&Reqdata)
		c.ProtoBuf(http.StatusOK, Resdata)

	} else {
		ginit.Scrlog.Error("VisitLogin Req data error")
		Resdata := &szprotobuf.ResLoginInfo{
			Code:4001,
			Msg:"error req data",
		}
		c.ProtoBuf(http.StatusOK, Resdata)
	}
}


func HandleFunc(Reqdata *szprotobuf.ReqVersion)(Resdata szprotobuf.ResVersionInfo){
	//循环break
	for {
		//数据校验
		if Reqdata.Plattype < 1 || Reqdata.Version < 1 {
			Resdata = szprotobuf.ResVersionInfo{
				Code:4001,
				Msg:"",
			}
			break
		}
		var VersionData szprotobuf.ResVersionData
		//停服公告检测
		result := models.CheckNotifition(Reqdata.Plattype, Reqdata.Gametype)
		if len( result) > 0{//需要停服公告
			var sysNotic models.ResNotification
			err := json.Unmarshal([]byte(result), &sysNotic)
			if err != nil{
				ginit.Scrlog.Error("Error Unmarshal err:%v", err)
			}
			VersionData.Notifi.Id = sysNotic.Id
			VersionData.Notifi.Title = sysNotic.Title
			VersionData.Notifi.Content = sysNotic.Content
			VersionData.Notifi.Starttime = sysNotic.Starttime
			VersionData.Notifi.Endtime = sysNotic.Endtime
			Resdata = szprotobuf.ResVersionInfo{
				Code: 200,
				Msg: "",
				Data: &VersionData,
			}
			break
		}
		//获取版本信息
		data := models.GetCurVersionInfo(Reqdata.Plattype, Reqdata.Gametype)
		var sysversion models.VersionInfo
		err := json.Unmarshal([]byte(data), &sysversion)
		if 	err != nil{
			ginit.Scrlog.Error("Error Unmarshal err:%v", err)
			Resdata = szprotobuf.ResVersionInfo{ //数据异常也返回不用更新
				Code: 200,
				Msg:  "",
			}
			break
		}
		//版本相同不用返回数据
		if Reqdata.Version == sysversion.CurVersion {
			Resdata = szprotobuf.ResVersionInfo{
				Code: 200,
				Msg:  "",
			}
			break
		}

		VersionData.Curversion = sysversion.CurVersion
		VersionData.Isforce = sysversion.IsForce
		VersionData.Updateurl = sysversion.UpdateUrl
		VersionData.Varmessage = sysversion.VarMessage
		VersionData.Vartitle =  sysversion.VarTitle
		Resdata = szprotobuf.ResVersionInfo{
			Code: 200,
			Msg: "",
			Data: &VersionData,
		}

		//跳出循环
		break
	}
	return Resdata
}
