package scribe

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"syscall"
)

type LocalFile struct {
	fi       *os.File
	filename string
	//bUpload  chan bool   //是否启动远程上传文件
	cFileLog chan string //本地日志记录
}

func GetHostName() string {
	host, err := os.Hostname()
	if err != nil {
		return "defaul"
	} else {
		return host
	}
}

func NewLocalFile(name string) (*LocalFile, error) {
	//bUpload := make(chan bool)
	cFileLog := make(chan string, 10)

	fi, err := os.OpenFile(name, syscall.O_RDWR|syscall.O_CREAT, 0666)
	if err != nil {
		fmt.Printf("[error]create file false %v\n", err)
		fi = nil
	}

	return &LocalFile{filename: name /*bUpload: bUpload,*/, cFileLog: cFileLog, fi: fi}, err
}

func (dofile *LocalFile) WriteFile(msg string) {
	if dofile.fi == nil {
		fi, err := os.Create(dofile.filename)
		if err != nil {
			fmt.Printf("[error]create file false %v msg:%v\n", err, msg)
			return
		}
		dofile.fi = fi
	}

	msg += "\r\n"
	_, err := dofile.fi.Write([]byte(msg))
	if err != nil {
		fmt.Printf("[warn] WriteFile %v", err)
		dofile.fi.Close()
		dofile.fi = nil
	}
}

func (dofile *LocalFile) UploadFile(doLog *GoScribe) {
	if dofile.fi != nil {
		dofile.fi.Close()
	}

	fi, err := os.Open(dofile.filename)
	if err != nil {
		fmt.Printf("[error]create file false %v \n", err)
		return
	}
	dofile.fi = fi

	br := bufio.NewReader(dofile.fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		//fmt.Println(string(a))
		if doLog.bBlockMode == true {
			doLog.cSendLog <- string(a)
		} else {
			doLog.ExeLog(string(a))
		}

	}

	dofile.fi.Close()
	os.Remove(dofile.filename)

	dofile.fi = nil
}

func (dofile *LocalFile) FileLoop(doLog *GoScribe) {
	fmt.Println("FileLoop .......")
	for {
		select {
		case msg := <-dofile.cFileLog:
			fmt.Println(msg)
			dofile.WriteFile(msg)

		}

	}
}
