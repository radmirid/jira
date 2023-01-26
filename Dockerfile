FROM golang:latest

ENV GO111MODULE=on

RUN go get github.com/andygrunwald/go-jira \
    && go get github.com/joho/godotenv \
    && go get github.com/lib/pq \
    && go get github.com/go-telegram-bot-api/telegram-bot-api

COPY . /app

WORKDIR /app

CMD ["go", "run", "main.go"]
