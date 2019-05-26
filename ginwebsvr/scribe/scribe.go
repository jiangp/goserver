package scribe

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/samuel/go-thrift/thrift"
)

/*
日志设置chan 模式  -- false 阻塞模式 （针对scribe）
doFile      本地日志
bBlockMode   bool //true 日志设置chan 模式  -- false 阻塞模式 （针对scribe）
bUserScribe  bool //true 使用scribe日志系统，false使用beego日志本地系统
bUnConScribe bool //true 正常连接scribe false无法连接scribe使用本地日志系统
hostName     string //服务器名，用于标记那台服务器上报日志
processName   string //程序进程名
remoteAddress string //scribe 连接地址

cSendLog chan string //bBlockMode:true 时候使用

mutex sync.Mutex //日志锁
*/
type GoScribe struct {
	log    *logs.BeeLogger
	scblog ScribeClient
	doFile *LocalFile //本地日志

	bBlockMode    bool        //true 日志设置chan 模式  -- false 阻塞模式 （针对scribe）
	bUserScribe   bool        //true 使用scribe日志系统，false使用beego日志本地系统
	bUnConScribe  bool        //true 正常连接scribe false无法连接scribe使用本地日志系统
	hostName      string      //服务器名，用于标记那台服务器上报日志
	processName   string      //程序进程名
	remoteAddress string      //scribe 连接地址
	cSendLog      chan string //bBlockMode:true 时候使用

	mutex sync.Mutex //日志锁
}

/*
*初始化日志
 */
func NewScribe(processName string, netAddr string, bUserScribe bool, bBlockMode bool) *GoScribe {
	hostName := GetHostName()
	err := os.Mkdir("log", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("create log ", err)
	}

	doLog := &GoScribe{processName: processName, remoteAddress: netAddr, hostName: hostName, cSendLog: make(chan string, 5000), bUserScribe: bUserScribe, bBlockMode: bBlockMode}
	//doLog.doFile, _ = NewLocalFile("log/" + processName + ".log")
	doLog.InitLog()
	return doLog
}

func (doLog *GoScribe) InitBeegoLog() {
	doLog.log = logs.NewLogger(100)
	//doLog.log.SetLogger("console", "")

	filedata := "{\"filename\":\"" + "log/" + doLog.processName + ".log" + "\"}"
	doLog.log.SetLogger("file", filedata)
	doLog.log.Async(1e3)
}

func (doLog *GoScribe) InitLog() {
	doLog.InitBeegoLog() //默认开启beegolog 日志库
	doLog.bUnConScribe = doLog.ReConnectScribe()
	if doLog.bUnConScribe {
		fmt.Printf("%v \n", "success connect scribe server !")
	}

	//if doLog.bBlockMode {
	go doLog.StartLog()
	//}

	//go doLog.doFile.FileLoop(doLog)
}

func (doLog *GoScribe) ReConnectScribe() bool {
	conn, err := net.Dial("tcp", doLog.remoteAddress)
	if err != nil {
		fmt.Println(err.Error())
		doLog.bUnConScribe = false
		return false
	}

	doLog.bUnConScribe = true
	t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(conn, 0), thrift.BinaryProtocol)

	client := thrift.NewClient(t, false)
	doLog.scblog = ScribeClient{Client: client}

	//doLog.doFile.bUpload <- true
	//go doLog.doFile.UploadFile(doLog)

	return doLog.bUnConScribe
}

func (doLog *GoScribe) SetLogHead(s string) string {

	pp, _, line, _ := runtime.Caller(2)

	if doLog.bUserScribe {
		now := time.Now()
		year, mon, day := now.Date()
		hour, min, sec := now.Clock()

		return fmt.Sprintf("[%v] %04d-%02d-%02d %02d:%02d:%02d [%v] [%v:%v] ",
			doLog.hostName, year, mon, day, hour, min, sec, s, runtime.FuncForPC(pp).Name(), line)

	} else {
		return fmt.Sprintf("[%v] [%v] [%v:%v] ",
			doLog.hostName, s, runtime.FuncForPC(pp).Name(), line)
	}
}

func (doLog *GoScribe) WriteLog(msg string) {
	if !doLog.bUnConScribe {
		if doLog.ReConnectScribe() == false {
			doLog.doFile.cFileLog <- msg
			return
		}
	}

	if len(msg) > 0 {
		_, err := doLog.scblog.Log([]*LogEntry{{doLog.processName, msg}})
		if err != nil {
			//doLog.log.Info(err.Error(), msg)
			//doLog.doFile.cFileLog <- msg
			doLog.bUnConScribe = false
			doLog.log.Info("%v \n error%v", msg, err)
		}

	}
}

func (doLog *GoScribe) StartLog() {
	doLog.Info("启动异步scribe 模式 ......")

	for {
		select {
		case msg := <-doLog.cSendLog:
			doLog.WriteLog(msg)
		}
	}
}

func (doLog *GoScribe) ExeLog(msg string) {
	//doLog.log.Info("%v", msg)
	if doLog.bUserScribe {

		if doLog.bBlockMode {
			doLog.cSendLog <- msg
			return
		}

		if !doLog.bUnConScribe {
			if doLog.ReConnectScribe() == false {
				doLog.log.Info("%v", msg)
				//doLog.doFile.cFileLog <- msg
				return
			}
		}

		_, err := doLog.scblog.Log([]*LogEntry{{doLog.processName, msg}})
		if err != nil {
			//doLog.log.Info(err.Error(), msg)
			doLog.bUnConScribe = false
			//doLog.doFile.cFileLog <- msg
			doLog.log.Info("%v", msg)
		}

	} else {
		doLog.log.Info("%v", msg)
	}

}

func (doLog *GoScribe) Info(format string, v ...interface{}) {
	msg := doLog.SetLogHead("Info") + fmt.Sprintf(format, v...)
	doLog.ExeLog(msg)
}

func (doLog *GoScribe) Warn(format string, v ...interface{}) {
	msg := doLog.SetLogHead("Warn") + fmt.Sprintf(format, v...)
	doLog.ExeLog(msg)
}

func (doLog *GoScribe) Error(format string, v ...interface{}) {
	msg := doLog.SetLogHead("Error") + fmt.Sprintf(format, v...)
	doLog.ExeLog(msg)
}

/*

func RunTest(doLog *GoScribe) {
	doLog.Info("test test %v", "1234556")
	time.Sleep(1)
}

func test() {
	doLog := NewScribe("category", "10.0.200.180:1463", GetHostName(), true, true)

	num := 1

	for i := 0; i < 20; i++ {
		go RunTest(doLog)
	}

	fmt.Printf("testestt .....\n")
	for {
		doLog.Info("do not come here %v", "tetetetet")
		num++
		time.Sleep(1 * time.Second)
	}

}
*/
