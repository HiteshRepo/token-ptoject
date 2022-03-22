FROM golang:1.17-alpine

WORKDIR /service

COPY . .

RUN go mod download

ENTRYPOINT ["go", "run", "cmd/server/main.go", "-configFile", "./config/default.yaml"]