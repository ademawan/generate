FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o load-testing-k6

EXPOSE 8098

CMD ./load-testing-k6