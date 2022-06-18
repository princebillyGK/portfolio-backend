FROM golang:1.18

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/ ./...


EXPOSE 8080
CMD ["api"]-

# Github package requirement
# LABEL org.opencontainers.image.source https://github.com/princebillyGK/my-portfolio-backend

