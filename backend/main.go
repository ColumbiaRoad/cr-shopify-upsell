package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/api"

	"github.com/labstack/gommon/log"
)

func main() {
	server := api.New()
	server.Router.Logger.SetLevel(log.DEBUG)

	server.Routes()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.Shutdown(5 * time.Second)

}
