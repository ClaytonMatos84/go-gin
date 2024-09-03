package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterNotFound(c *gin.Context)  {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
