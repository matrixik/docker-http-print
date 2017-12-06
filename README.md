# Docker HTTP print

Start it on some port (default to "8080") and it will print all request data
that will reach it.

You can change default port with `PORT` env variable.

```bash
$ docker build -t docker-http-print .
$ docker run -e PORT=8090 -it --rm -p 8090:8090 docker-http-print
```

In Docker Compose file:
```yaml
  http-print:
    image: docker-http-print
    environment:
      PORT: 8090
      LOGSPOUT: ignore
    ports:
      - "8090:8090"
```

## From Docker Hub:

```bash
$ docker run -e PORT=8090 -it --rm -p 8090:8090 matrixik/http-print
```

In Docker Compose file:
```yaml
  http-print:
    image: matrixik/http-print
    environment:
      PORT: 8090
      LOGSPOUT: ignore
    ports:
      - "8090:8090"
```

