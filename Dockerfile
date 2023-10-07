FROM golang:1.18-alpine

WORKDIR /app

COPY . .
RUN  go build  -o ./bin/crawller-ofsistorage ./cmd/main.go
#COPY ${HOME}/.postgresql/root.crt /usr/local/share/ca-certificates/
#ENTRYPOINT ["go", "test", "-v", "./...", "-coverprofile", "cover.out"]
ENTRYPOINT ["./bin/crawller-ofsistorage"]