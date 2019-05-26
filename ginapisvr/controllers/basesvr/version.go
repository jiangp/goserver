package basesvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Req
type ReqVersion struct {
	LoginType uint32 `form:"loginType" json:"loginType" binding:"required"` //1-FB, 2-游客
	Type      uint32 `form:"type" json:"type" binding:"required"`           // 1-安卓 2-ios ...
	Version   uint32 `form:"version" json:"version" binding:"required"`     //版本号
}

//Res
type ResVersion struct {
	models.ResHeader
	Data ResData `form:"data" json:"data" `
}

type ResData struct {
	CurVersion uint32 `form:"curVersion" json:"curVersion" `
	VarTitle   string `form:"varTitle" json:"varTitle" `
	VarMessage string `form:"varMessage" json:"varMessage" `
	IsForce    uint32 `form:"isForce" json:"isForce" `
	UpdateUrl  string `form:"updateUrl" json:"updateUrl" `
}

func CheckVsion(c *gin.Context) {
	ginit.Scrlog.Info("Req %v", c.Request.Body)
	var ReqData ReqVersion
	if err := c.ShouldBindJSON(&ReqData); err == nil {
		//数据校验
		if ReqData.Type < 1 || ReqData.Version < 1 {
			c.JSON(http.StatusOK, gin.H{
				"code": 4001, //错误的请求数据
				"msg":  "Error Request Data",
			})
			return
		}
		if ReqData.Type == 1 { //安卓版本检测
			Resbody := &ResVersion{
				ResHeader: models.ResHeader{Code: 200, Msg: ""},
				Data: ResData{
					CurVersion: 1000,
					VarTitle:   "新春特别版！",
					VarMessage: "安卓平台新年特惠",
					IsForce:    0,
					UpdateUrl:  "www.baidu.com",
				},
			}
			c.JSON(http.StatusOK, Resbody)
			return

		} else if ReqData.Type == 2 { //ios版本检测
			Resbody := &ResVersion{
				ResHeader: models.ResHeader{Code: 200, Msg: ""},
				Data: ResData{
					CurVersion: 1000,
					VarTitle:   "新春特别版！",
					VarMessage: "IOS平台新年特惠",
					IsForce:    0,
					UpdateUrl:  "www.baidu.com",
				},
			}
			c.JSON(http.StatusOK, Resbody)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 4001,
			"msg":  "Error Request Data",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 4001,
			"msg":  "Error Request Data",
		})
	}
}
