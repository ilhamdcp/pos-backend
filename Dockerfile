FROM golang:1.18
WORKDIR .
COPY src .
RUN ["go", "run", "src/main.go"]