FROM golang:1.19-alpine as builder
# FROM 172.18.241.238 as builder

# mkdir app
# cd app
WORKDIR /app

COPY ./ ./

RUN go mod tidy

RUN go build -o ./bin/cms ./main.go

# ------------ buat (build stage) baru------
FROM alpine:3
ENV APP_NODE=1

WORKDIR /app

# copy env dan binary file
COPY --from=builder /app/.env ./
COPY --from=builder /app/bin/cms ./

EXPOSE 5000

CMD ./cms