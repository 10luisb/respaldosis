FROM golang

WORKDIR /grpcClient

COPY . .

RUN go mod tidy
RUN go mod download

EXPOSE 3000

CMD ["go", "run", "Main.go"]