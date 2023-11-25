FROM golang:1.21.3
WORKDIR /app
RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/sidsreeram/ecom.git

COPY /cmd/db.env /app/ecom/cmd
WORKDIR /app/ecom/cmd
ENTRYPOINT [ "go","run","main.go"]