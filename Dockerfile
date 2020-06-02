FROM golang:1.14

COPY . /app
WORKDIR /app
RUN go build -o starwars .

ENTRYPOINT ["/app/starwars"]