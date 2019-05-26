package main

import (
	"ginwebsvr/ginit"
	_ "ginwebsvr/ginit"
	"ginwebsvr/routers"
)

func main() {
	//开始启动
	addr := ginit.Conf.DefaultString("port", "4080")
	routers.R.Run(addr) // listen and serve on 0.0.0.0:4080
}
