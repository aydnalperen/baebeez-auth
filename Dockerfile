FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go mod tidy
RUN go build -o /dist

EXPOSE 8080

CMD [ "/dist" ]