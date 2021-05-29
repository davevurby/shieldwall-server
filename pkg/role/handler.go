package role

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/davevurby/shieldwall"
	"github.com/davevurby/shieldwall-server/core/http/helpers"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	store shieldwall.Store
}

func NewRoleHandler(store shieldwall.Store) *RoleHandler {
	return &RoleHandler{store: store}
}

// @Router /v1/roles/{id} [put]
func (h *RoleHandler) PutRole(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	var body = struct {
		Namespaces []string `json:"namespaces"`
	}{}

	err := json.Unmarshal(bodyBytes, &body)
	if err != nil {
		helpers.AbortNamespacesInvalid(c)
		return
	}

	role := shieldwall.Role{Id: c.Param("id"), Namespaces: body.Namespaces}
	err = h.store.PutRole(role)
	if err != nil {
		log.Fatalf("unable to store role - %s\n", err.Error())
	}

	c.AbortWithStatusJSON(201, role)
}
