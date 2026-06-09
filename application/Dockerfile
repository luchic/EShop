FROM golang:1.26.1

RUN mkdir /app

COPY . /app

WORKDIR /app/server

EXPOSE 8080

CMD ["go", "run", "./cmd/shop/main.go"]