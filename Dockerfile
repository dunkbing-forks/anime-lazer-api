# BUILD
FROM golang:1.18-alpine as BUILD
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api
ENV HTTP_PORT=3000
EXPOSE $HTTP_PORT

# BINARIES
FROM alpine:latest
COPY --from=BUILD /app/api /app/api
ENTRYPOINT ["/app/api"]
