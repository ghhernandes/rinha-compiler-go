FROM golang:1.21.1

WORKDIR /var/app

RUN apt update && apt upgrade -y

RUN git clone https://github.com/ghhernandes/rinha-compiler-go.git /var/app

RUN GOOS=linux go build -o rinha /var/app/cmd/main.go

RUN cp rinha /usr/local/bin

ENTRYPOINT [ "/bin/bash" ]
