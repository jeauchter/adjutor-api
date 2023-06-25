FROM golang:latest
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY *.go .
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
CMD "air"