BUILD_TIME=$(shell date +%FT%T%z)
VERSION=`git describe --tags`

build_backend:
	@GOOS=linux build -ldflags "-w -s -X main.buildTime=${BUILD_TIME} -X main.version=${VERSION}" -o backend

deploy_backend: build_backend
	@docker build -t razeencheng/ssltools_backend:${VERSION} .
	@docker push razeencheng/ssltools_backend:${VERSION}

