package drivers

import (
	"github.com/gin-gonic/gin"
	"github.com/zumak/zumo/backend/store"
)

var drivers struct {
	servers map[string]ServerInitFunc
	store   map[string]StoreInitFunc
}

func init() {
	drivers.servers = map[string]ServerInitFunc{}
	drivers.store = map[string]StoreInitFunc{}
}

type ServerInitFunc func(gin.IRouter)

func ServerRegister(name string, driver ServerInitFunc) {
	drivers.servers[name] = driver
}
