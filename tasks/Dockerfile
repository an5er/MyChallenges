FROM golang:1.16

RUN apt-get update && apt-get install -y sqlite3

WORKDIR /go/src/
RUN mkdir -p github.com/thewhitetulip/

WORKDIR /go/src/github.com/thewhitetulip/
RUN git clone https://github.com/thewhitetulip/Tasks.git

WORKDIR /go/src/github.com/thewhitetulip/Tasks

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct

RUN go mod init
RUN go get github.com/dgrijalva/jwt-go@v3.0.0
RUN go get github.com/gorilla/sessions@v1.1

RUN apt-get install -y pkg-config libsqlite3-dev
RUN go get github.com/mattn/go-sqlite3@v1.4.0
RUN go get github.com/shurcooL/github_flavored_markdown
RUN go build -o tasks
EXPOSE 8081

COPY schema.sql /go/src/github.com/thewhitetulip/Tasks/
RUN cat schema.sql | sqlite3 tasks.db

CMD ["/go/src/github.com/thewhitetulip/Tasks/tasks"]