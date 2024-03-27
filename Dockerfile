FROM golang:latest as builder
LABEL maintainer="Rajkumar <rajkumaar2304@icloud.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy -e
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
RUN if [ "$TARGETARCH" == "amd64"] ; \
    then curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64 ; \
    else curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-arm64 ; \
    fi
RUN chmod +x /usr/local/bin/dbmate
WORKDIR /app/
COPY --from=builder /app/main .
COPY --from=builder /app/db/ .
EXPOSE 2610
CMD /usr/local/bin/dbmate up && ./main