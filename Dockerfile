FROM golang:alpine

ENV APP_PATH=$GOPATH/src/app
RUN mkdir -p $APP_PATH/frontend

WORKDIR $APP_PATH

COPY go.mod  go.sum ./
RUN go mod download

COPY frontend/package.json frontend/yarn.lock ./frontend/

RUN apk update && apk add npm nodejs build-base && npm -g i yarn && cd frontend/ && yarn install;

ADD . ./
RUN cd frontend/; yarn build;

EXPOSE 9000
EXPOSE 9090

RUN go build -o /usr/local/bin/jpay ./cmd/main.go && chmod +x /usr/local/bin/jpay

CMD ["/usr/local/bin/jpay", "-dsn", "$APP_PATH/testdb/sample.db", "-limit", "15"]




