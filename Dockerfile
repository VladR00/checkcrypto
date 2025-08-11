FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod tidy
COPY . .
RUN go build -o /app/builds/app /app/cmd/main.go

FROM alpine
WORKDIR /
COPY --from=builder /app/builds/app builds/app
ENV CONFIGPATH=configs/dev.yaml
CMD ["./builds/app"]
EXPOSE 8080