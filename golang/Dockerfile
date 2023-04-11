FROM golang:1.19 as builder
WORKDIR /build
COPY go.mod  .
RUN go mod download
COPY . . 
RUN GOOS=linux GOARCH=amd64 go build -o /build/entrypoint

FROM scratch
COPY --from=builder /build/entrypoint /go/build/entrypoint
ENTRYPOINT ["/go/build/entrypoint"]