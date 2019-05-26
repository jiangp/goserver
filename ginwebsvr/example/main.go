package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"szprotobuf"
	"time"

	"github.com/gogo/protobuf/proto"
)

const (
	LHST  = "http://:4080/"
	LHST1 = "http://132.232.228.201/"
	LHST2 = "http://lite.workec.com/"
)

func main() {
	httpproto()
	//httpjson()

}

func httpproto() {
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
	prodata := szprotobuf.ReqLoginInfo{
		Userid: 0,
		Faceid: "2271076426552352",
		Token:  "EAAIy3UpEPpIBAJCAq1g3LfpbzgkuUEXKBYLKrwIBrzmZBNwMqtg9uTAvGjHXliQzeyn0BQIlsXGCPGAAGqPjsmKEuG60MaEFg2NPZCA5Q4OgFaplNZBZBxNJFeU4VRZBhNrYG6PWK6E57JQH4qTIkZA1OVZCfVDxBaZB9buCH5hSndH6q9RuVmMTpkdUA1eDBnx4ocaz6uIlWlMOOyGXRNZC9nOuuKZCPEbEYZD",
		//Imei:   "xxxxxxxxxxx",
	}

	data, err := proto.Marshal(&prodata)

	req, err := http.NewRequest("POST", LHST+"login/v1/visit", bytes.NewBuffer(data))
	//req, err := http.NewRequest("POST", LHST+"version/v1/update", bytes.NewBuffer(data))
	req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("X-Ec-Pvkey", "34W0iMLkTadrtC4zmc1Ik1GRBf6pbUpeNdP9aKEIwMs=@@")
	//req.Header.Set("X-Ec-Uid", "5009386")
	req.Header.Set("Content-Type", "application/x-protobuf")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func httpjson() {
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
	var jsonStr = []byte(`{"loginType":2,"type":1, "version":1111}`)

	//req, err := http.NewRequest("POST", LHST+"login/v1/visit", bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest("POST", LHST1+"version/v1/update", bytes.NewBuffer(jsonStr))
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
