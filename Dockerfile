FROM golang:1.12

LABEL maintainer="Fajar"

WORKDIR $GOPATH/src/github.com/farkroft/gin_rest_tutorial

COPY . .

RUN make install-dep

RUN make setup

RUN make build

RUN ls -la

EXPOSE 80

CMD ["./rest_gin_tutorial"]