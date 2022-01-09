package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"num":  1,
		"data": nil,
	})
}
