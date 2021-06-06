FROM golang

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV PORT 6000

COPY . /go/src/app

WORKDIR /go/src/app

RUN go get ./

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o servicemysql

CMD ./servicemysql

EXPOSE 6000