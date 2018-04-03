package server

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"

	"github.com/zumak/zumo/backend"
)

const (
	keyUsername = "USERNAME"
)

// Run is
func Run(conf *Config, b backend.Backend) error {
	app := gin.Default()

	server := &Server{rice.MustFindBox("../dist"), b}

	app.StaticFS("/static", server.dist.HTTPBox())

	app.GET("/", server.CheckAuth, server.static("text/html", "html/index.html"))
	app.GET("/register", server.static("text/html", "html/register.html"))
	app.POST("/register", server.register)

	app.GET("/ws", server.CheckAuth, server.ws)

	v1 := app.Group("/api/v1", server.CheckAuth)
	{
		v1.GET("/channels", server.getChannels)
		v1.POST("/channels", server.createChannel)
		//v1.DELETE("/channel", server.deleteChannel)

		v1.PUT("/channels/:channelID/join", server.joinChannel)
		v1.PUT("/channels/:channelID/invite/:username", server.invite)
		v1.PUT("/channels/:channelID/kick/:username", server.kick)

		//v1.GET("/user/:username", server.getUserInfo)
		v1.GET("/users/:username/joinned-channel", server.joinnedChannel)

		// Messages
		v1.GET("/channels/:channelID/messages", server.getMessage)
		v1.POST("/channels/:channelID/messages", server.postMessage)
	}

	return app.Run(conf.Bind)
}

// Server is
type Server struct {
	dist    *rice.Box
	backend backend.Backend
}

// Config is
type Config struct {
	Bind string
}

func (server *Server) static(contextType, file string) func(c *gin.Context) {
	buf := server.dist.MustBytes(file)
	return func(c *gin.Context) {
		c.Data(http.StatusOK, contextType, buf)
	}
}
