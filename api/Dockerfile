# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

# Download modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

# copy everything to /app
COPY . ./

# build
RUN go build -o /docker-webCli

EXPOSE 3000

CMD [ "/docker-webCli" ]


