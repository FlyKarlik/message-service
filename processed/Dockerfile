FROM golang:1.22-alpine3.20 AS builder 


RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM alpine AS production 
WORKDIR /app
COPY --from=builder /app/server .


EXPOSE 3000
CMD [ "./server" ]