FROM golang:latest as builder
LABEL maintainer="Rajkumar <rajkumaar2304@icloud.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy -e
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/main .
EXPOSE 2610
CMD ["./main"]