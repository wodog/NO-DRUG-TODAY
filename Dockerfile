FROM golang

RUN go get github.com/wodog/NO-DRUG-TODAY

WORKDIR /go/src/github.com/wodog/NO-DRUG-TODAY

CMD go run main.go
