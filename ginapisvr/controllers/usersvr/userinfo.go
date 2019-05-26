package usersvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqUserInfo struct {
	UserId int64  `form:"userid" json:"userid" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
	//Data   []byte `form:"data" json:"data"`
}

func GetUserInfo(c *gin.Context) {
	ginit.Scrlog.Info("Req:%v", c.Request)
	var ReqData ReqUserInfo
	if err := c.ShouldBindJSON(&ReqData); err == nil {
		//数据校验
		if ReqData.UserId < 1000 || len(ReqData.Token) < 20 {
			c.JSON(http.StatusOK, gin.H{
				"code": 4001, //错误的请求数据
				"msg":  "Error Request Data",
			})
			return
		}

		//鉴权
		err := utils.VerifyKey(ReqData.UserId, ReqData.Token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 401, //鉴权失败
				"msg":  "Error Request Data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK, //鉴权失败
			"msg":  "",
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 4001,
			"msg":  "Error Request Data",
		})
	}
}
