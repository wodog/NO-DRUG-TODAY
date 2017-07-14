FROM golang

RUN go get github.com/wodog/NO-DRUG-TODAY

WORKDIR /go/src/github.com/wodog/NO-DRUG-TODAY

RUN git config --global user.name wodog \
&& git config --global user.email 536505032@qq.com \
&& git remote set-url origin git@github.com:wodog/NO-DRUG-TODAY

CMD go run main.go
