package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	LHST  = "http://:4080/"
	LHST1 = "http://192.168.1.70:8088/"
	LHST2 = "http://lite.workec.com/"
)

func main() {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(25 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*20)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	var jsonStr = []byte(`{"message":"xxxxxx","nick":"nihao", "userid":1111}`)

	req, err := http.NewRequest("POST", LHST+"login/v1/visit", bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("X-Ec-Pvkey", "34W0iMLkTadrtC4zmc1Ik1GRBf6pbUpeNdP9aKEIwMs=@@")
	//req.Header.Set("X-Ec-Uid", "5009386")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
