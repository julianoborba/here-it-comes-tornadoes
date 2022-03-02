FROM golang:1.17 as build-env
WORKDIR /go/src/app
COPY main.go .
COPY go.sum ./
COPY go.mod ./
COPY docs/ ./docs/
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=build-env /go/bin/app /
CMD ["/app"]
