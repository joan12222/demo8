FROM golang:1.17 AS build
WORKDIR /demo8
COPY . .
ENV CGO_ENABLED=0
ENV GO11MODULE=on
ENV GOOS=linux go build -installsuffix cgo -o bin/httpserver main.go

FROM busybox
COPY --from=build /demo8/bin/httpserver /demo8/bin/httpserver
EXPOSE 8360
WORKDIR /demo8/
ENTRYPOINT ["./bin/httpserver"]