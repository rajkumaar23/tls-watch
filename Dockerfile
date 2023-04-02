FROM golang:1.20.2-alpine3.17

RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN go build -o ssl-expiry-watcher

RUN echo "23 4 * * * cd /app && ./ssl-expiry-watcher" >> /var/spool/cron/crontabs/root
CMD crond -f