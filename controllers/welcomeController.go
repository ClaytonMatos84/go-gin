package controllers

import "github.com/gin-gonic/gin"

func Salutation(c *gin.Context)  {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"Seja bem-vindo": name + ". Tudo bom?",
	})
}
