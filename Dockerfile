FROM golang:1.20 as build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/adjutor

#final
FROM alpine:latest
COPY --from=build /opt/adjutor /opt/adjutor
EXPOSE 8080
ENTRYPOINT [ "/opt/adjutor" ]

