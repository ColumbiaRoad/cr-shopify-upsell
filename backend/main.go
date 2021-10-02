package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/db"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/app/merchant"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/api"

	"github.com/labstack/gommon/log"
)

func main() {
	apiSecret, found := os.LookupEnv("SHOPIFY_API_SECRET")
	if !found {
		log.Fatalf("missing env variable SHOPIFY_API_SECRET")
	}
	apiKey, found := os.LookupEnv("SHOPIFY_API_KEY")
	if !found {
		log.Fatalf("missing env variable SHOPIFY_API_KEY")
	}
	backendURL, found := os.LookupEnv("BACKEND_URL")
	if !found {
		log.Fatalf("missing env variable BACKEND_URL")
	}

	redirectURL := backendURL + "/v1/shopify/callback"
	server := api.New(apiKey, apiSecret, redirectURL)
	server.Router.Logger.SetLevel(log.DEBUG)

	server.Routes()
	db, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}
	m := merchant.New(db)
	server.Merchant = m
	if err != nil {
		panic(err)
	}

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
