FROM golang:1.22 AS build
LABEL authors="p1mako"

WORKDIR /cart-api

COPY . .

RUN go mod download && go mod verify

RUN go build -v -o /cart-api/build/cart-api cmd/cart-api/cart-api.go

FROM ubuntu:22.04

COPY --from=build /cart-api/build/cart-api /bin/cart-api
COPY --from=build /cart-api/internal/config/db-conf.json /internal/config/db-conf.json
EXPOSE 3000

ENTRYPOINT ["/bin/cart-api"]
