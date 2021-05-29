package v1

import "github.com/gin-gonic/gin"

func GetHealth(c *gin.Context) {
	c.AbortWithStatus(200)
}
