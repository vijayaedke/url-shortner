FROM golang:1.19.2-alpine as builder

WORKDIR /opt/url-shortner

ENV CGO_ENABLED=0 
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8086

ENTRYPOINT [ "/opt/cmd/urlapp/main" ]
