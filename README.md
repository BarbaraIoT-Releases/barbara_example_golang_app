This is a Barbara OS example Golang app.

It needs a basic appConfig to work:

```
{
    "boolKey": true,
    "stringKey": "foo",
    "intKey": 5,
    "debugLevel": 3
}
```

To build this app, Docker and Docker Compose are needed.

Modify "platform" inside docker-compose.yaml to "linux/amd64", "linux/arm64" or "linux/arm/v7" depending on your platform.

To build the app, run:
```
docker compose up --build
```

Built app will be in "output" folder.