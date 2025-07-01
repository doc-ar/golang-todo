FROM golang:1.24.4

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

RUN go build -o /usr/local/bin/app ./cmd

EXPOSE 8080

CMD ["app"]
