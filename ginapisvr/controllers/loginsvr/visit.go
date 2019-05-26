package loginsvr

import (
	"ginwebsvr/ginit"
	"ginwebsvr/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID  int64  `form:"userid" json:"userid" `
	Nick    string `form:"nick" json:"nick" `
	Faceid  int64  `form:"faceid" json:"faceid" `
	FbToken string `form:"facetoken" json:"facetoken"`
}

func VisitLogin(c *gin.Context) {
	ginit.Scrlog.Info("Req:%v", c.Request)
	var ReqData User
	if err := c.ShouldBindJSON(&ReqData); err == nil {
		//生成uid
		if ReqData.UserID < 1000 && ReqData.Faceid < 1 {
			ReqData.UserID = utils.GetNewUserID()
		}

		//生成key
		Token := utils.GetUserKey(ReqData.UserID)
		if len(Token) < 1 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Tips": "I'm going to have coffee. Please have a rest.",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":    "",
			"token":  Token,
			"userid": ReqData.UserID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 4001,
			"msg":  "error req data",
		})
	}

}
