FROM golang:1.20 as builder

WORKDIR /app
COPY . .

RUN go build -v -o server ./server

FROM golang:1.20

COPY --from=builder /app/server /server

CMD ["/server"]
