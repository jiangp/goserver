package ginit

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"os"
)


var (
	Conf      config.Configer
)

func InitConf() {
	var err error
	Conf, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("read config failed", err)
		os.Exit(-1)
	}
}