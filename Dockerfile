FROM golang:1.20-buster as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build \
  -tags timetzdata \
  -o /go/bin/app \
  cmd/chat/main.go cmd/chat/wire_gen.go


FROM gcr.io/distroless/static-debian11 as production

COPY --from=builder /go/bin/app /

EXPOSE 8080

CMD ["/app"]