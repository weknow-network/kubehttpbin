# dockerized development environment variables
REPO_PATH := github.com/arschles/kubehttpbin
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.9.1
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -e GO15VENDOREXPERIMENT=1 -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}

VERSION ?= git-$(shell git rev-parse --short HEAD)
IMAGE := quay.io/arschles/kubehttpbin:${VERSION}

LDFLAGS := "-s -X main.version=${VERSION}"

bootstrap:
	${DEV_ENV_CMD} glide install

build:
	${DEV_ENV_PREFIX} -e CGO_ENABLED=0 ${DEV_ENV_IMAGE} go build -a -installsuffix cgo -ldflags ${LDFLAGS} -o rootfs/bin/kubehttpbin

test:
	${DEV_ENV_CMD} sh -c 'go test $$(glide nv)'

docker-build:
	docker build --rm -t ${IMAGE} rootfs

docker-push:
	docker push ${IMAGE}
