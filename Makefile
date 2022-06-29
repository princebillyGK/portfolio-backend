include .env
export

run: 
	cd frontend && npm run build
	cd cmd/api/ && go run main.go
deploy:
	docker build -f Dockerfile.ci -t ghcr.io/princebillygk/portfolio-backend:latest . 

