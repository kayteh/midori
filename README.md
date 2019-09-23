# Midori

Devmerald's ChatOps Discord Bot. 

## Using

### Docker

This project is published to both `katie/midori` and `docker.pkg.github.com/kayteh/midori/midori`. Both include AMD64, ARM64v8, and ARM32v7 images.

Using it would look like:
```sh
docker run -it -p3030:3030 -v/path/to/certs/:/etc/certs/ \
    -e GITHUB_CLIENT_SECRET=... \
    -e GITHUB_APP_CERT_PATH=... \
    -e DISCORD_BOT_TOKEN=... \
        katie/midori
```

## Contributing

Pre-requisites

- Go 1.13+ (probably earlier too) w/ `GO111MODULE=on`
- Discord Bot Token -- [from here](https://discordapp.com/developers)
- GitHub App w/ .pem -- [from here](https://github.com/settings/apps/new)

After clone
```
go get ./...
go run .
go test ./...
```

Yay!

### Building
