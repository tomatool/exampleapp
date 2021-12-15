FROM golang:1.17

COPY . /app

WORKDIR /app

ENTRYPOINT [ "go" ]
CMD ["run", "main.go"]