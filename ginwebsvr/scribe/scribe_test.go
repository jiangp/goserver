package scribe

import (
	"fmt"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	doLog := NewScribe("testlog", "10.0.200.180:1463", true, true)

	num := 1
	/*
		for i := 0; i < 20; i++ {
			go RunTest(doLog)
		}
	*/
	fmt.Printf("testestt .....\n")
	for {
		doLog.Info("do not come here %v", "tetetetet")
		num++
		time.Sleep(1 * time.Second)
	}
}
