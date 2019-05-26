package ginit

import (
	"git.workec.com/scribe"
	"os"
)

var (
	Scrlog *scribe.GoScribe
)

func InitScrlog() error {
	InitScribeLog("ginwebsvr")
	return nil
}

func InitScribeLog(serviceName string) {
	if serviceName == "" {
		os.Exit(-1)
	}

	Scrlog = scribe.NewScribe(serviceName,
		"127.0.0.1",
		false,
		true,
	)
	Scrlog.Info("init scribe log success")
}