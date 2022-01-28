FROM golang:1.16-alpine

WORKDIR /app
ENV GOPROXY https://goproxy.cn,direct
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8000

CMD [ "/docker-gs-ping" ]