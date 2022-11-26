########-- Build stage --########

FROM golang:1.19.3-alpine AS builder-prod

RUN  apk add git

WORKDIR /opt

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . ./
RUN go build -o /myapp ./cmd/rest


########-- Deploy stage --########

FROM alpine:3.16

WORKDIR /opt 

COPY --from=builder-prod /myapp /opt/myapp
COPY ./.env .

EXPOSE 8080

CMD [ "/opt/myapp"]