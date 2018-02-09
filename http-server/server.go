package server

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

const (
	keyUsername = "USERNAME"
)

// Run is
func Run(conf *Config) error {
	app := gin.Default()

	server := &Server{rice.MustFindBox("../dist")}

	app.StaticFS("/static", server.dist.HTTPBox())

	return app.Run(conf.Bind)
}

// Server is
type Server struct {
	dist *rice.Box
}

// Config is
type Config struct {
	Bind string
}
