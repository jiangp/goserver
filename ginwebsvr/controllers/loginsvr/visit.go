package loginsvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/models"
	"ginwebsvr/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"szprotobuf"
	"time"
)


func VisitLogin(c *gin.Context) {
	var Reqdata szprotobuf.ReqLoginInfo
	if err := c.ShouldBind(&Reqdata); err == nil{
		ginit.Scrlog.Info("VisitLogin Req:%v ficeid:%v", Reqdata.Userid, Reqdata.Faceid)
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

func HandleFunc(Reqdata *szprotobuf.ReqLoginInfo)(Resdata szprotobuf.ResLoginInfo) {
	//生成uid
	for{
		//设备信息校验
		if len(Reqdata.Imei) < 1{
			Resdata = szprotobuf.ResLoginInfo{
				Code:4001,
				Msg:"用户设备硬件码错误！",
			}
			break
		}
		var LoginInfo szprotobuf.ResLoginData

		//访客
		if Reqdata.Userid < 1000 && Reqdata.Faceid < 1 {
			//生成新用户
			Reqdata.Userid = utils.GetNewUserID()
			err:= models.SaveUserInfo(Reqdata.Userid, Reqdata.Nickname)
			if err != nil{
				Resdata = szprotobuf.ResLoginInfo{
					Code:1,
					Msg:"systeam error!",
				}
				break
			}
		}else if Reqdata.Userid < 1000 && Reqdata.Faceid > 0 {//第三方登录
			if len(Reqdata.Token) < 1{
				Resdata = szprotobuf.ResLoginInfo{
					Code:4001,
					Msg:"请求数据错误！",
				}
				break
			}
			//第三方登录校验
			err := utils.OtherLoginVerify(Reqdata.Othertype,Reqdata.Faceid, Reqdata.Token)
			if err != nil{
				Resdata = szprotobuf.ResLoginInfo{
					Code:4002,//检验失败
					Msg:"第三方校验不通过",
				}
				break
			}
			//查询对应用户用户基本信息//名称和uid
			models.SearchUserInfo(Reqdata.Faceid)
		}

		//生成key
		Token := utils.GetUserKey(Reqdata.Userid)
		if len(Token) < 1 {
			Resdata = szprotobuf.ResLoginInfo{
				Code:1,
				Msg:"I'm going to have coffee. Please have a rest.",
			}
			break
		}
		//登录信息
		LoginInfo.Nowtime = time.Now().Unix()
		LoginInfo.Userid = Reqdata.Userid
		LoginInfo.Nickname = Reqdata.Nickname
		LoginInfo.Svrtoken = Token
		Resdata = szprotobuf.ResLoginInfo{
			Code:200,
			Msg:"",
			Data:&LoginInfo,
		}

		break
	}

	return Resdata
}
