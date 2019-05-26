package utils

import (
	"errors"
	"fmt"
	"ginwebsvr/ginit"
	"time"
)

//获取自增uid
func GetNewUserID() int64 {
	Result, err := ginit.GetRedis().Incr("Const_Global_UID").Result()
	if err != nil {
		return 0
	}
	return Result
}

//目前每次登录获取新的token导致旧的失效
func GetUserKey(UserID int64) string {
	var Token string
	Key := fmt.Sprintf("User_Key:%v", UserID)
	//Result , err := ginit.GetRedis().Get(Key).Result()
	//if err != nil { //没有已经存在的key

	//生成用户Key
	Token = fmt.Sprintf("%v_%v", string(Krand(16, KC_RAND_KIND_ALL)), string(Krand(8, KC_RAND_KIND_ALL)))
	_, err := ginit.GetRedis().Set(Key, Token, 24*time.Hour).Result() //有效期3天
	if err != nil {
		ginit.Scrlog.Error("err:%v", err)
		time.Sleep(1 * time.Second)
		_, err = ginit.GetRedis().Set(Key, Token, 24*time.Hour).Result()
		if err != nil {
			return Token // 服务异常报错
		}
	}
	//}else{
	//	Token =  Result  //保证状态， 可以去除
	//}
	return Token

}

func VerifyKey(UserID int64, Token string) error {
	Key := fmt.Sprint("User_Key:%v", UserID)
	Result, err := ginit.GetRedis().Get(Key).Result()
	if err != nil { //没有已经存在的key
		return errors.New("Error User Auth")
	} else {
		if Result != Token {
			return errors.New("Error User Token")
		}
		//添加有效期维护逻辑
		ginit.GetRedis().Expire(Key, 24*time.Hour)
	}
	return nil
}
