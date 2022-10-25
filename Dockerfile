# Dockerfile.production

FROM golang:1.19-bullseye as builder

ENV APP_HOME /go/src/demoapp

WORKDIR "$APP_HOME"
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o demoapp

FROM golang:1.19-bullseye

ENV APP_HOME /go/src/demoapp
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY src/conf/ conf/
COPY src/views/ views/
COPY --from=builder "$APP_HOME"/demoapp $APP_HOME

EXPOSE 80
CMD ["./demoapp"]
