package main

import (
	"github.com/zumak/zumo"
	// server
	//_ "github.com/zumak/server/core"
	// store drivers
	_ "github.com/zumak/zumo/backend/store/etcd"
)

func main() {
	zumo.Run()
}
