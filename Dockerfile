FROM golang:1.24.2 AS dev

WORKDIR /workspace

COPY . .

CMD ["go", "run", "./app/cmd/main.go"]


FROM golang:1.24.2 AS builder

WORKDIR /workspace

COPY ./app ./app
COPY ./go.mod ./go.sum ./

RUN go mod tidy

RUN go build -o binary ./app/cmd/main.go


FROM gcr.io/distroless/base AS prod

WORKDIR /workspace

COPY --from=builder --chown=nonroot:nonroot /workspace/binary ./binary

CMD ["./binary"]