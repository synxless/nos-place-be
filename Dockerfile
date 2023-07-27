FROM golang:1.20-alpine3.16 AS build

RUN apk update && apk add gcc musl-dev gcompat libc-dev linux-headers
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tc-token *.go

FROM alpine:3.16
EXPOSE 8100

WORKDIR /app

COPY --from=build /app/tc-token /app/tc-token

RUN chmod +x /app/tc-token
ENTRYPOINT [ "/app/tc-token" ]
