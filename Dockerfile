FROM golang:latest
WORKDIR /app
COPY ./example /app
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build example.go" --command=./example
