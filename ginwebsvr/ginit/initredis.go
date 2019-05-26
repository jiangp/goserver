package ginit

import (
	"time"

	"gopkg.in/redis.v5"
)

var (
	redisclient *redis.Client //redis base
)

func InitRedis() error {
	var err error
	//基础redis
	addr := Conf.DefaultString("redis::addr", "127.0.0.1:6379")
	pwd := Conf.DefaultString("redis::password", "12345678")
	redisclient, err = NewRedis(addr, pwd, 10, time.Second*3)
	if err != nil {
		Scrlog.Info("Init Redis addr:%v %v", addr, err)
		return err
	}
	//redis ...
	return nil

}

func NewRedis(addr string, password string, PoolSize int, DialTimeout time.Duration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		PoolSize:    PoolSize,
		DialTimeout: DialTimeout,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, err
}

func GetRedis() *redis.Client {
	return redisclient
}