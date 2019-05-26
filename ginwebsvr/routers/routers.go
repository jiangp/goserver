package routers

import (
	"ginwebsvr/controllers"
	"ginwebsvr/controllers/loginsvr"
	"ginwebsvr/controllers/usersvr"
	"ginwebsvr/controllers/basesvr"
	"github.com/gin-gonic/gin"
	"net/http"
)


var (
	R *gin.Engine
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(404)
	rw.Write([]byte("page_not_found"))

}

func init(){
	R = gin.Default()
	//R.GET("404", page_not_found)

	R.GET("/ping", controllers.Pingfunc)
	//版本检测
	version := R.Group("/version")
	{
		version.POST("/v1/update", basesvr.CheckVsion)
	}

	//登录
	login := R.Group("/login")
	{
		login.POST("/v1/visit", loginsvr.VisitLogin)
		//login.POST("/v1/ficebook", loginsvr.VisitLogin)
		//login.POST("/v1/google", loginsvr.VisitLogin)
	}

	//用户信息
	user := R.Group("/user")
	{
		user.POST("/v1/userinfo", usersvr.GetUserInfo)
		user.POST("/v1/infochange", usersvr.GetUserInfo)
		user.POST("/v1/nickchange", usersvr.GetUserInfo)
		user.POST("/v1/battlerecord", usersvr.GetUserInfo)//战绩列表
	}

	//好友信息
	friend := R.Group("/friend")
	{
		friend.POST("/v1/friendlist", controllers.PostTest)
		friend.POST("/v1/searchfriend", controllers.PostTest)
		friend.POST("/v1/addfriend", controllers.PostTest)
		friend.POST("/v1/delfriend", controllers.PostTest)
		friend.POST("/v1/aceptfriend", controllers.PostTest)
	}

	//背包
	knapsack := R.Group("/knapsack")
	{
		knapsack.POST("/v1/proplist", controllers.PostTest)
		knapsack.POST("/v1/propinfo", controllers.PostTest)
		knapsack.POST("/v1/propuser", controllers.PostTest)
	}

	//商城
	mall := R.Group("/mall")
	{
		mall.POST("/v1/malllist", controllers.PostTest)
		//mall.POST("/v1/pay", controllers.PostTest)
		mall.POST("/v1/recharge", controllers.PostTest)//支付校验
		mall.POST("/v1/rushbag", controllers.PostTest)//首冲礼包
	}

	//广播
	broadcast := R.Group("/broadcast")
	{
		broadcast.POST("/v1/sendbroad", controllers.PostTest)
		//broadcast.POST("/v1/ficebook", controllers.PostTest)  //接收广播svr推送
		//broadcast.POST("/v1/google", controllers.PostTest)
	}

	//邮件
	email := R.Group("/email")
	{
		email.POST("/v1/emaillist", controllers.PostTest)
		email.POST("/v1/emailmsg", controllers.PostTest)//午评获取
		//email.POST("/v1/google", controllers.PostTest) //在线邮件推送
	}

	//观看视屏
	video := R.Group("/video")
	{
		video.POST("/v1/vidoetoken", controllers.PostTest)
		video.POST("/v1/vidoebonus", controllers.PostTest)
	}

	//评价校验
	evaluate := R.Group("/evaluate")
	{
		evaluate.POST("/v1/evaluatecheck", controllers.PostTest)
		//evaluate.POST("/v1/evaluatecheck", controllers.PostTest)
	}

	//排行榜
	rankinglist := R.Group("/ranking")
	{
		rankinglist.POST("/v1/goldcoin", controllers.PostTest)
		rankinglist.POST("/v1/integral", controllers.PostTest)
	}

}