package main

import (
	"context"
	"log"

	"princebillygk.portfolio.io/conf"
	"princebillygk.portfolio.io/repository"
)

func main() {
	ctx := context.TODO()
	mc, dbClose, err := conf.NewMongoClient(ctx)
	if err != nil {
		log.Fatalln("Failed to connect to MongoDB")
	}
	defer dbClose()

	portfolioRepo := repository.NewPortfolio(mc)
	portfolioRepo.GetById(ctx, conf.ResumeId)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("Hello World"))
	// })

	// http.ListenAndServe("0.0.0.0:"+util.Env("API_PORT", "80"), nil)
}
