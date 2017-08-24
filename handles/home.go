package handles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeHandle struct {
}

func InitHomeHandle(e *gin.Engine) {
	h := HomeHandle{}

	// Setup Routes
	e.GET("/", h.homePage)
}

func (h *HomeHandle) homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{})
}
