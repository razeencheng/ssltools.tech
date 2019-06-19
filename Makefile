BUILD_TIME=$(shell date +%FT%T%z)
VERSION=`git describe --tags`

build_backend:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X main.buildTime=${BUILD_TIME} -X main.version=${VERSION}" -o backend

deploy_backend: build_backend
	@docker build -t razeencheng/ssltools_backend:${VERSION} .
	@docker push razeencheng/ssltools_backend:${VERSION}

deploy_nginx:
	@docker build -t razeencheng/nginx_tls13:1.7.0_1.1.1 ./nginx
	@docker push razeencheng/nginx_tls13:1.7.0_1.1.1