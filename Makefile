APP=skeleton
MAIN=cmd/main.go
BINARY=bin/${APP}
PKG:=github.com/iamasmith/${APP}
VP:=${PKG}/internal/version
# TODO: Autogenerate off tag etc.
LATEST_TAG:=$(shell git describe --abbrev=0 HEAD)
RELEASE_TAG=$(shell git tag --points-at HEAD)
CURRENT_SHA=$(shell git rev-parse HEAD)
VERSION:=${LATEST_TAG}
ifeq ($(LATEST_TAG),$(RELEASE_TAG))
BUILDTYPE:=release
else
BUILDTYPE:=snapshot-${CURRENT_SHA}
endif
ifeq ($(BUILDID),)
BUILDTYPE:=$(shell whoami)-workstation-${CURRENT_SHA}
BUILDID:=None
endif
BUILDDATE:=$(shell date -u)

all: test build

test:
	go test --cover "./internal/..."

build:
	CGO_ENABLED=0 go build -o ${BINARY} -ldflags '-X "${VP}.LName=${APP}" -X "${VP}.LVersion=${VERSION}" -X "${VP}.LBuildDate=${BUILDDATE}" -X "${VP}.LBuildType=${BUILDTYPE}" -X "${VP}.LBuildId=${BUILDID}"' ${MAIN}

clean:
	go clean
	rm ${BINARY}
