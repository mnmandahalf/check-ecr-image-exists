FROM golang:1.19.2-alpine

ENV GOPATH=

COPY . .

RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
