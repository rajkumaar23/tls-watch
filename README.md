# ssl expiry watch 

[![watch cron](https://github.com/rajkumaar23/ssl-expiry-watch/actions/workflows/cron.yml/badge.svg)](https://github.com/rajkumaar23/ssl-expiry-watch/actions/workflows/cron.yml)

> a lightweight tool that monitors SSL certificate expirations, providing timely alerts through Telegram, with easy configuration via environment variables.

## prerequisites
ensure you have `go` installed.

## configuration
set the following environment variables, either in a .env file (or) your pipeline configuration:
- `HOSTNAMES`: Comma-separated list of hostnames.
- `THRESHOLD`: Integer threshold.
- `TELEGRAM_ID`: Telegram user ID.
- `TELEGRAM_BOT_TOKEN`: Telegram bot token.

## usage
1) build the application:
```bash
go build -o ssl-expiry-watch
```
2) run the executable:
```bash
./ssl-expiry-watch
```

## license
this project is licensed under the [MIT License](//github.com/rajkumaar23/ssl-expiry-watch/tree/main/LICENSE).
