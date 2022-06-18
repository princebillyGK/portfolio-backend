run:
	go run ./cmd/api/main.go
deploy:
	docker build -f Dockerfile -t ghcr.io/princebillygk/portfolio-backend:latest . 

