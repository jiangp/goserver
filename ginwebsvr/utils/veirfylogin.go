package utils

import (
	"fmt"
	"ginwebsvr/ginit"
	"github.com/astaxie/beego/httplib"
	"time"
)

var (
	//fb
	AccessUrl = "https://graph.facebook.com/oauth/access_token"
	CheckUrl = "https://graph.facebook.com/debug_token"
	AppID = "618875968568978"
	SecretID = "4a5fe807559bfbb498528fb1f0698782"
)

//第三方登录校验
func OtherLoginVerify(Type uint32, Ficeid string, Token string) error{
	if Type == 0{ //FB登录检验
		return FbTokenCheck(Ficeid, Token)
	}
	return nil
}

func FbTokenCheck(Ficeid string, Token string) error{
	httpReq := CheckUrl + fmt.Sprintf("?access_token={%v}|{%v}&input_token=%v",AppID,SecretID ,Token)
	req := httplib.Get(httpReq).SetTimeout(5* time.Second, 3 * time.Second)
	str, err := req.String()
	if err != nil{
		ginit.Scrlog.Error(" data :%v err:%v", str, err)
	}

	//req.Response()
	ginit.Scrlog.Info("%v", str)
	return  nil
}