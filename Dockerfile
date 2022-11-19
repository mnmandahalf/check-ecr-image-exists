FROM golang:1.19.2-alpine

RUN mkdir /app
WORKDIR /app

COPY . .

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
