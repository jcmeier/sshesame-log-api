FROM golang as build-env
WORKDIR /go/src/sshesame-log-api
ADD . /go/src/sshesame-log-api
RUN CGO_ENABLED=0 go build -o /go/bin/sshesame-log-api
FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/sshesame-log-api /
CMD ["/sshesame-log-api"]
