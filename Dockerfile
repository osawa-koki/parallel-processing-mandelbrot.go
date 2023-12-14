FROM golang:1.20
WORKDIR /work/
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
CMD ["go", "test", "-bench", "."]
