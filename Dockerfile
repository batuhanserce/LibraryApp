FROM golang:1.16-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o temp

EXPOSE 3000

ENTRYPOINT [ "./temp" ]