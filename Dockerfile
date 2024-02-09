
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/ecom-exe ./cmd/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/ecom-exe /app/

COPY template /app/template

COPY cmd/db.env /app/db.env

RUN chmod +x /app/ecom-exe

EXPOSE 3000

CMD ["/app/ecom-exe"]
