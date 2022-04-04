.DEFAULT_GOAL := docker-image

IMAGE ?= cncamp/demo8-server:latest

image/demo8-server: $(shell find . -name '*.go')
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o $@ ./cmd/main.go

.PHONY: docker-image
docker-image: image/demo8-server
	docker build -t $(IMAGE) image/

.PHONY: push-image
push-image: docker-image
	docker push $(IMAGE)
