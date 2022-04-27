FROM golang:alpine

ENV APP_PATH  $GOPATH/src/app

RUN mkdir -p $APP_PATH/app/frontend

COPY go.mod $APP_PATH/

COPY go.sum $APP_PATH/

WORKDIR $APP_PATH

RUN go mod download

COPY frontend/package.json $APP_PATH/frontend/

COPY frontend/yarn.lock $APP_PATH/frontend/

RUN apk update  && apk add npm nodejs build-base

RUN npm -g i yarn

RUN cd $APP_PATH/frontend && yarn install;

ADD . $APP_PATH/

RUN cd $APP_PATH/frontend; yarn build;

EXPOSE 9000

EXPOSE 9090

RUN go build -o /usr/local/bin/jpay $APP_PATH/cmd/main.go

CMD ["/usr/local/bin/jpay"]




