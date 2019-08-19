CONTAINER_VERSION=latest

.PHONY: image

image: working_server broken_server
	docker build -t working_server:${CONTAINER_VERSION} -f Dockerfile .

working_server: working-server.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' working-server.go

broken_server: broken-server.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' broken-server.go

