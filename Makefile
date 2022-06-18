run:
	go run ./cmd/api/main.go
deploy:
	docker build -t ghcr.io/princebillygk/my-portfolio-backend:latest .
	docker push ghcr.io/princebillygk/portfolio-backend


