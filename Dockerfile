FROM golang:1.10

WORKDIR /opt/sleep-server

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]