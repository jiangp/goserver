package models

import (
	"fmt"
	"ginapisvr/ginit"
)



func CheckNotifition(PlanType uint32, GameType uint32) string {
	Result, err := ginit.GetRedis().HGet("CheckNotifition",fmt.Sprintf("key:%v:%v:1", GameType,PlanType)).Result()
	if err != nil {
		ginit.Scrlog.Error("GameCurVersion:%v err :%v", PlanType, err)
		return ""
	}

	return Result
}
func GetCurVersionInfo(PlanType uint32, GameType uint32) string {

	Result, err := ginit.GetRedis().Get(fmt.Sprintf("GameCurVersion:%v:%v", GameType,PlanType)).Result()
	if err != nil {
		ginit.Scrlog.Error("GameCurVersion:%v err :%v", PlanType, err)
		return ""
	}

	return Result

}
