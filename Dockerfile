FROM golang:1.12
LABEL maintainer="scottedwardrussell@gmail.com"
WORKDIR /build
COPY . /build
RUN go mod download && go test -v && CGO_ENABLED=0 GOOS=linux go build -x -installsuffix cgo -o go-example-app .

FROM alpine:latest
WORKDIR /app
COPY --from=0 /build/go-example-app .
EXPOSE 8000
ENTRYPOINT ["./go-example-app"]
