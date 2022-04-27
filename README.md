## Jpay Exercise
> A Simple Golang application built using Gin Framework + Element Plus + Vite and Vue3

This is a simple Go application that fetches data from an SQL database and exposes a rest API and a `gRPC` service on port `localhost:9090`
Technology stack
 - Backend  (Go with Gin framework) 
 - Frontend (VuJS & element-ui)
 - Database (SQLite3)

### Getting Started
Clone the repository to your local directory and use docker to deploy. this will expose the application on `http://localhost:9000`. Once the page loads ypo can use the dropdown to filter results based on the listed countries.

```shell
git clone https://github.com/weezqyd/jpay.git

cd jpay && docker compose up --build
```
## Development

### Backend
We use nodemon to watch for file changes and recompile the go application everytime we make a change the `.go` files.
```shell
yarn install
#Run server and watch file changes
yarn run dev ./cmd/main.go

```
### Frontend
Set up the frontend for local development

```bash
cd frontend
yarn install
yarn run dev
```


### Running Tests
To run the included e2e and unit tests, open another terminal tab and run the command below.

```shell
docker exec -it jpay-exercise-dev go test -v ./internal/...
```
