package basesvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/logic"
	"net/http"
	"szprotobuf"

	"github.com/gin-gonic/gin"
)

func CheckVsion(c *gin.Context) {

	var Reqdata szprotobuf.ReqVersion
	if err := c.ShouldBind(&Reqdata); err == nil{
		ginit.Scrlog.Info("VisitLogin Logintype:%v Type:%v, Version:%v", Reqdata.Logintype, Reqdata.Plattype, Reqdata.Version)
		Resdata := logic.VersionCheck(&Reqdata)
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

