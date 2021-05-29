package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davevurby/shieldwall-server/core"
	"github.com/davevurby/shieldwall-server/core/config"
	"github.com/davevurby/shieldwall-server/core/http/router"
)

func main() {
	cfg := config.GetConfig()
	store := core.CreateStore(cfg)

	routesHandler := router.CreateRouter(store)
	port := cfg.GetHttpPort()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: routesHandler,
	}

	log.Printf("Server is going to be started on port %d\n", port)
	log.Fatal(server.ListenAndServe())
}
