FROM golang:1.18-alpine as builder

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go
FROM alpine:latest
WORKDIR /app  
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata

COPY .env /app/
COPY --from=builder /app/main .

EXPOSE 8186
CMD ["/app/main"]