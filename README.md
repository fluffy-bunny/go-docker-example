# go-docker-example

[reference](https://www.bacancytechnology.com/blog/dockerize-golang-application)  
[golang-docker-example](https://github.com/DekivadiyaKishan/golang-docker-example)  

This is an example of how to run 2 web servers in docker container.  Each listening on different ports.

## Step By Step Build and Run

### Prerequisites  

* [Docker](https://docs.docker.com/engine/install/)
* [golang](https://go.dev/dl/)

### Clone

```bash
git clone https://github.com/fluffy-bunny/go-docker-example.git
```

### Build locally

In a console navigate to where you cloned the project.  

```bash
cd go-docker-example
go mod tidy
go build ./...
go test ./...
```

### Build Docker Image

```bash
docker build --rm -t golang-docker-example .
```

### Run via Docker-Compose

```bash
docker-compose up
```

### Stopping Docker-Compose

```bash
docker-compose down
```
