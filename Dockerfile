FROM golang:latest AS build

RUN go version

ENV GOPATH=/

# Copying all project files
COPY . .

# Installing psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# making wait-for-postgres.sh executable via bash script
RUN chmod +x wait-for-postgres.sh

# building go app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o socket_server ./cmd/app/main.go

CMD ["./server"]
