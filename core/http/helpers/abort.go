package helpers

import "github.com/gin-gonic/gin"

func AbortNamespacesInvalid(c *gin.Context) {
	c.AbortWithStatusJSON(400, struct {
		Message string `json:"message"`
	}{
		Message: "namespaces are invalid",
	})
}
