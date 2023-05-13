APP=skeleton
MAIN=cmd/main.go
BINARY=bin/${APP}
PKG:=github.com/iamasmith/${APP}
VP:=${PKG}/internal/version
# TODO: Autogenerate off tag etc.
VERSION:=v0.0.1
BUILDID:=00000000
BUILDTYPE:=workstation-$(shell whoami)
BUILDDATE:=$(shell date -u)

all: test build

test:
	go test --cover "./internal/..."

build:
	CGO_ENABLED=0 go build -o ${BINARY} -ldflags '-X "${VP}.LName=${APP}" -X "${VP}.LVersion=${VERSION}" -X "${VP}.LBuildDate=${BUILDDATE}" -X "${VP}.LBuildType=${BUILDTYPE}" -X "${VP}.LBuildId=${BUILDID}"' ${MAIN}

clean:
	go clean
	rm ${BINARY}
