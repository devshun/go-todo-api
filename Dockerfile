FROM golang:alpine3.15

WORKDIR /app/src

COPY ./main.go .

RUN go mod init todo && go mod tidy

CMD ["go", "run", "./main.go"]