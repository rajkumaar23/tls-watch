# ssl-expiry-watch

To run ssl-expiry-watch on your system, follow the steps below :

- Ensure you have `docker` and `docker-compose` installed.
- Create a directory for storing the config and docker-compose files.
- Get inside *that* directory.
- Create `config.json` with the contents resembling the format of this [example config file](https://github.com/rajkumaar23/ssl-expiry-watch/blob/main/config.example.json).
- Create `docker-compose.yml` inside the same directory with the content below :
  - **DO NOT FORGET** to update the `/absolute/path/to/config.json` inside the compose file
```yaml
version: "3"

services:
  ssl-expiry-watch:
    image: ghcr.io/rajkumaar23/ssl-expiry-watch:main
    container_name: ssl-expiry-watch
    restart: unless-stopped
    volumes:
      - /absolute/path/to/config.json:/app/config.json
```
- Spin up the container with `docker-compose up -d` and enjoy!
