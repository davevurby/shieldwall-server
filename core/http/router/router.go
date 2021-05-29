package router

import (
	"github.com/davevurby/shieldwall"
	v1 "github.com/davevurby/shieldwall-server/core/http/router/v1"
	"github.com/davevurby/shieldwall-server/pkg/role"
	"github.com/gin-gonic/gin"
)

func CreateRouter(store shieldwall.Store) *gin.Engine {
	roleHandler := role.NewRoleHandler(store)

	r := gin.New()

	r.GET("/__health", v1.GetHealth)

	r.GET("/v1/roles/{id}", roleHandler.PutRole)

	return r
}
