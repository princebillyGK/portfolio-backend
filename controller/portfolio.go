package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"princebillygk.portfolio.io/conf"
	"princebillygk.portfolio.io/repository"
)

type Portfolio struct {
	repo *repository.Portfolio
}

func NewPortfolio(dbClient *mongo.Client) Portfolio {
	return Portfolio{
		repo: repository.NewPortfolio(dbClient),
	}
}

func (p Portfolio) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := p.repo.GetById(ctx, conf.ResumeId)
	if err != nil {
		log.Fatalln(err)
	}

	s, _ := json.MarshalIndent(data, "", "\t")
	log.Println(string(s))

	w.WriteHeader(http.StatusOK)
	conf.Templates["homepage"].Execute(w, data)
}
