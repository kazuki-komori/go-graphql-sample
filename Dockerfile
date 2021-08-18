FROM golang:1.15.7-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=off

WORKDIR ${ROOT}

RUN apk update && apk add git
COPY . ${ROOT}

RUN go mod download

EXPOSE 8080

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/99designs/gqlgen

CMD ["go", "run", "server.go"]