package logic

import "szprotobuf"

func UserInfo(Reqdata *szprotobuf.ReqUserInfo) (Resdata szprotobuf.ResUserInfo) {
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