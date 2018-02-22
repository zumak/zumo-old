package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckAuth is
func (server *Server) CheckAuth(c *gin.Context) {
	str := c.GetHeader("Authorization")
	if str == "" {
		// Request Auth
		c.Header("WWW-Authenticate", "Basic realm=\"Auth required!\"")
		c.Data(http.StatusUnauthorized, "text/html", server.dist.MustBytes("html/unauthorized.html"))
		c.Abort()
		return
	}

	token, err := server.backend.Token(str)
	if err != nil {
		c.String(http.StatusUnauthorized, "Token Not Found: %s", err.Error())
		c.Abort()
		return
	}
	c.Set(keyUsername, token.Username)

	return // continue to next
}

func (server *Server) register(c *gin.Context) {
	req := &struct {
		ID       string
		Password string
	}{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	if _, err := server.backend.CreateUser(req.ID, map[string]string{"zumo.type": "user"}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	if _, err := server.backend.CreateToken(req.ID, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}
