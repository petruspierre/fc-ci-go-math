FROM golang:1.15 as builder

WORKDIR /app

COPY . .

RUN go build -o math

FROM scratch

COPY --from=builder /app/math .

CMD ["./math"]