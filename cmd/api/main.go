package main

import (
	"context"
	"log"
	"net/http"

	"princebillygk.portfolio.io/conf"
	"princebillygk.portfolio.io/controller"
	"princebillygk.portfolio.io/pkg/util"
	"princebillygk.portfolio.io/router"
)

func main() {
	ctx := context.TODO()
	mc, dbClose, err := conf.NewMongoClient(ctx)
	if err != nil {
		log.Fatalln("Failed to connect to MongoDB")
	}
	defer dbClose()

	// Setuping up router
	r := router.New()
	router.SetupPortfolio(r, controller.NewPortfolio(mc))

	// Serving static files
	r.Handle("/*", http.FileServer(http.Dir(conf.StaticAssetsPath)))

	log.Fatal(http.ListenAndServe("0.0.0.0:"+util.Env("API_PORT", "80"), r))
}
