package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"princebillygk.portfolio.io/conf"
	"princebillygk.portfolio.io/pkg/util"
	"princebillygk.portfolio.io/repository"
)

func main() {
	ctx := context.TODO()
	mc, dbClose, err := conf.NewMongoClient(ctx)
	if err != nil {
		log.Fatalln("Failed to connect to MongoDB")
	}
	defer dbClose()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		portfolioRepo := repository.NewPortfolio(mc)
		data, err := portfolioRepo.GetById(ctx, conf.ResumeId)
		if err != nil {
			log.Fatalln(err)
		}
		s, _ := json.MarshalIndent(data, "", "\t")
		w.Write([]byte(s))
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:"+util.Env("API_PORT", "80"), nil))
}
