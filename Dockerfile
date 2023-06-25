FROM golang:1.20 as build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/rest-service

#final
FROM alpine:latest
COPY --from=build /opt/rest-service /opt/rest-service
EXPOSE 8080
ENTRYPOINT [ "/opt/rest-service" ]
