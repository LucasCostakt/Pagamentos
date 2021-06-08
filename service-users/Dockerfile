FROM golang

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV PORT 5050

COPY . /go/src/app

WORKDIR /go/src/app

RUN go get ./

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o serviceusers

CMD ./serviceusers

EXPOSE 5050