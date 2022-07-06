FROM golang:latest as builder
WORKDIR /app
ENV GOPROXY https://goproxy.io
COPY . .
RUN go mod tidy -compat=1.17
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN mkdir -p /app/config
WORKDIR /app
COPY --from=builder /app .
ENTRYPOINT ["./main"]
