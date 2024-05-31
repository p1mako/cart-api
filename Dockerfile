FROM golang:1.22
LABEL authors="p1mako"

COPY . /cart-api

WORKDIR /cart-api

RUN go mod download && go mod verify

RUN go build -v -o build/cart-api cmd/cart-api/cart-api.go

EXPOSE 3000

ENTRYPOINT /cart-api/build/cart-api
