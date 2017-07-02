FROM golang

RUN go get github.com/wodog/NO-DRUG-TODAY

WORKDIR src/github.com/wodog/NO-DRUG-TODAY

CMD go run main.go
