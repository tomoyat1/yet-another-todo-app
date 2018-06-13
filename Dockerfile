FROM golang:1.10-alpine

WORKDIR /go/src/github.com/tomoyat1/yet-another-todo-app

RUN apk add --update curl git make netcat-openbsd \
	&& rm -rf /var/cache/apk/*

COPY ./ ./
RUN make install

EXPOSE 8080
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["/go/bin/todo-server"]
