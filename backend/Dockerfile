FROM golang:1.22.3

RUN apt-get update && apt-get install -y make

WORKDIR /backend

COPY  go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/backend ./cmd/api/main.go

EXPOSE 3345

RUN chmod +x ./bin

CMD ["./bin/backend"]