FROM golang:alpine
ENV APP_PATH  $GOPATH/src/app
RUN mkdir -p $APP_PATH/frontend
WORKDIR $APP_PATH
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY frontend/package.json ./frontend/
COPY frontend/yarn.lock ./frontend/
RUN apk update && apk add npm nodejs build-base
RUN npm -g i yarn
RUN cd frontend/ && yarn install;
ADD . .
RUN cd frontend/; yarn build;
EXPOSE 9000
EXPOSE 9090
RUN go build -o /usr/local/bin/jpay ./cmd/main.go
CMD ["/usr/local/bin/jpay"]




