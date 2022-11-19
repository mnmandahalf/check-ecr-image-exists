FROM golang:1.19.2-alpine

COPY . .

RUN ["chmod", "+x", "./entrypoint.sh"]

ENTRYPOINT ["./entrypoint.sh"]
