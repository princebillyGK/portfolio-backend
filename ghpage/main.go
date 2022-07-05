package main

import (
	"context"
	"log"
	"os"

	"princebillygk.portfolio.io/conf"
	"princebillygk.portfolio.io/repository"
)

const outputPath = `../../frontend/public/index.html`

func main() {
	ctx := context.TODO()
	mc, dbClose, err := conf.NewMongoClient(ctx)
	if err != nil {
		log.Fatalln("Failed to connect to MongoDB")
	}
	defer dbClose()

	repo := repository.NewPortfolio(mc)
	data, err := repo.GetById(ctx, conf.ResumeId)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = conf.Templates["homepage"].Execute(f, data)
	if err != nil {
		log.Fatalln(err)
	}
}
