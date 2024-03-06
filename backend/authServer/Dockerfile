FROM golang:1.21.4-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cashe add bash gcc gettext

# dependenices
COPY ["./../go.mod", "./../go.sum", "./"]
RUN go mod download

#build
COPY . ./
RUN go build -o ./bin/app cmd/main.go

FROM alpine

COPY --from=builder /usr/local/src/bin/app /
COPY configs/config.yaml /config.yaml

CMD ["/app"]