FROM golang:1.13.4-alpine3.10

WORKDIR /go/src/slack-repeat-bot
COPY . .

RUN go build

WORKDIR /opt
RUN mv /go/src/slack-repeat-bot/slack-repeat-bot .
RUN rm -rf /go/src/slack-repeat-bot

ENTRYPOINT ["/opt/slack-repeat-bot"]
