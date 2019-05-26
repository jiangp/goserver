package logic

import (
	"ginwebsvr/models"
	"ginwebsvr/utils"
	"szprotobuf"
	"time"
)

func Loginverify(Reqdata *szprotobuf.ReqLoginInfo) (Resdata szprotobuf.ResLoginInfo) {
	//生成uid
	for {
		//设备信息校验
		if len(Reqdata.Imei) < 1 {
			Resdata = szprotobuf.ResLoginInfo{
				Code: 4001,
				Msg:  "用户设备硬件码错误",
				Data: nil,
			}
			break
		}
		var LoginInfo szprotobuf.ResLoginData

		//访客
		if Reqdata.Userid < 1000 && len(Reqdata.Faceid) > 1 {
			//生成新用户
			Reqdata.Userid = utils.GetNewUserID()
			err := models.SaveUserInfo(Reqdata.Userid, Reqdata.Nickname)
			if err != nil {
				Resdata = szprotobuf.ResLoginInfo{
					Code: 1,
					Msg:  "systeam error!",
					Data: nil,
				}
				break
			}
		} else if len(Reqdata.Faceid) > 1 { //第三方登录
			if len(Reqdata.Token) < 1 {
				Resdata = szprotobuf.ResLoginInfo{
					Code: 4001,
					Msg:  "请求数据错误！",
					Data: nil,
				}
				break
			}
			//第三方登录校验
			err := utils.OtherLoginVerify(Reqdata.Othertype, Reqdata.Faceid, Reqdata.Token)
			if err != nil {
				Resdata = szprotobuf.ResLoginInfo{
					Code: 4002, //检验失败
					Msg:  "第三方校验不通过",
					Data: nil,
				}
				break
			}
			//与系统内部帐号进行匹配
			//无对应用户
			//有对应用户

			//查询对应用户用户基本信息//名称和uid
			models.SearchUserInfo(Reqdata.Faceid)
		}

		//生成key
		Token := utils.GetUserKey(Reqdata.Userid)
		if len(Token) < 1 {
			Resdata = szprotobuf.ResLoginInfo{
				Code: 1,
				Msg:  "I'm going to have coffee. Please have a rest.",
				Data: nil,
			}
			break
		}
		//登录信息
		LoginInfo.Nowtime = time.Now().Unix()
		LoginInfo.Userid = Reqdata.Userid
		LoginInfo.Nickname = Reqdata.Nickname
		LoginInfo.Svrtoken = Token
		Resdata = szprotobuf.ResLoginInfo{
			Code: 200,
			Msg:  "",
			Data: &LoginInfo,
		}

		break
	}

	return Resdata
}
