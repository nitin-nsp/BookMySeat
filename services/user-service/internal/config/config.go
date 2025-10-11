package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Postgres struct {
		Host         string
		Port         int
		User         string
		Password     string
		Database     string
		MaxOpenConns int
		MaxIdleConns int
	}
	Redis struct {
		Host string
		Type string
		Pass string
	}
	JWT struct {
		Secret string
		Expire int64
	}
}
