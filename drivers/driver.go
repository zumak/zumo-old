package drivers

import "github.com/gin-gonic/gin"

var drivers struct {
	servers map[string]ServerInitFunc
}

func init() {
	drivers.servers = map[string]ServerInitFunc{}
}

type ServerInitFunc func(gin.IRouter)

func ServerRegister(name string, driver ServerInitFunc) {
}
