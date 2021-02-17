# make docker command overridable (eg for using `sudo docker` or `podman` instead)
DOCKER_COMMAND ?= docker

REPO := quay.io/routerd/go-dhcp-playground
TAG := $(shell date +"%s")

.PHONY: release build push

release: | build push

build:
	docker build -t "$(REPO):$(TAG)" .

push:
	docker push "$(REPO):$(TAG)"