package loginsvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"szprotobuf"
)


func VisitLogin(c *gin.Context) {
	var Reqdata szprotobuf.ReqLoginInfo
	if err := c.ShouldBind(&Reqdata); err == nil{
		ginit.Scrlog.Info("VisitLogin Req:%v ficeid:%v", Reqdata.Userid, Reqdata.Faceid)
		Resdata := logic.Loginverify(&Reqdata)
		c.ProtoBuf(http.StatusOK, Resdata)

	} else {
		ginit.Scrlog.Error("VisitLogin Req data error")
		Resdata := &szprotobuf.ResVersionInfo{
			Code:4001,
			Msg:"error req data",
		}
		c.ProtoBuf(http.StatusOK, Resdata)
	}

}

