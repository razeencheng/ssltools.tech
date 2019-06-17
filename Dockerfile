FROM alpine

LABEL maintainer="NGINX Docker Maintainers <me@razeen.me>"

RUN apk update && apk add ca-certificates
ADD backend /app/backend

WORKDIR /app

EXPOSE 8080

CMD ["./backend"]
