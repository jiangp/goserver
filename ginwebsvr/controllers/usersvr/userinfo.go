package usersvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/utils"
	"net/http"
	"szprotobuf"

	"github.com/gin-gonic/gin"
)

type ReqUserInfo struct {
	UserId int64  `form:"userid" json:"userid" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
	//Data   []byte `form:"data" json:"data"`
}

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

		Resdata := HandleFunc(&Reqdata)
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

func HandleFunc(Reqdata *szprotobuf.ReqUserInfo) (Resdata szprotobuf.ResUserInfo) {
	for{
		//查询用户信息
		Resdata = szprotobuf.ResUserInfo{
			Code:200,
			Msg:"",
		}
		break
	}
	return Resdata
}
