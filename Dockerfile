FROM golang:1.18-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main-app

EXPOSE 3000

ENTRYPOINT [ "./main-app" ]