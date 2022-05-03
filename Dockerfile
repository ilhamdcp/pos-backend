FROM golang:1.18

WORKDIR ~/pos-backend

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

CMD ["go", "run", "main.go"]