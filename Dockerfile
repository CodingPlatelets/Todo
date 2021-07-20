FROM golang:latest

ENV GO111MODULE=on \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /build

COPY . .
RUN go build Todo

EXPOSE 5001

CMD ["/build/Todo"]

