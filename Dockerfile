FROM golang:1.21.1

WORKDIR /var/app

RUN apt update && apt upgrade -y && apt install -y

RUN git clone https://github.com/ghhernandes/rinha-compiler-go.git /var/app

RUN CGO_ENABLED=0 GOOS=linux go build -o rinha /var/app/cmd/main.go

ENTRYPOINT [ "/bin/bash" ]
