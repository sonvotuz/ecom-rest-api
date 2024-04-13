FROM golang:1.22.1 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out ./cmd/main.go

# Build the Go app
RUN go build -o out ./cmd/main.go

# Run the application on container startup
CMD ["./out"]