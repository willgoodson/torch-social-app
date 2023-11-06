FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o echo-server .

EXPOSE 1323

CMD ["./echo-server"]