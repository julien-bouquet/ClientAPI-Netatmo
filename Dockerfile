# build stage
FROM golang:1.13-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

# final stage
FROM scratch
COPY --from=builder /app/ /app/
ENTRYPOINT [ "netatmo-client" ]