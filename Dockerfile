FROM golang:1.24.2 AS dev

WORKDIR /workspace

COPY . .

CMD ["go", "run", "./app/cmd/main.go"]


FROM golang:1.24.2 AS builder

WORKDIR /app

COPY ./app .
COPY ./go.mod ./go.sum .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o binary ./cmd/main.go


FROM gcr.io/distroless/base AS prod

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /app/binary ./binary

CMD ["./binary"]