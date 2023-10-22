# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /opt/url-shortner

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build cmd/urlapp/main.go

EXPOSE 8086
ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
CMD ["./main"]