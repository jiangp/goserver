package usersvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/logic"
	"ginwebsvr/utils"
	"net/http"
	"szprotobuf"

	"github.com/gin-gonic/gin"
)


func GetUserInfo(c *gin.Context){
	var Reqdata szprotobuf.ReqUserInfo
	if err := c.ShouldBind(&Reqdata); err == nil{
		ginit.Scrlog.Info("VisitLogin Req:%v Token:%v", Reqdata.Userid, Reqdata.Svrtoken)
		//登录校验
		err = utils.VerifyKey(Reqdata.Userid, Reqdata.Svrtoken)
		if err != nil {
			Resdata := szprotobuf.ResLoginInfo{
				Code:401,
				Msg:"err key！",
			}
			c.ProtoBuf(http.StatusOK, Resdata)
			return
		}

		Resdata := logic.UserInfo(&Reqdata)
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


