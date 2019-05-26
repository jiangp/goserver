package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Message string `form:"message" json:"message" binding:"required"`
	Nick string `form:"nick" json:"nick" binding:"required"`
}

func PostTest(c *gin.Context){
	fmt.Println(c.Request)
	var json User
	if err := c.ShouldBindJSON(&json); err == nil{
		message := json.Message
		nick := json.Nick

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	}

}