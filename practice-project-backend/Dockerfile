FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base && apk cache clean
ENV CGO_ENABLED=1
RUN go build -o ./proj-backend ./cmd/proj-backend/main.go


FROM alpine
WORKDIR /app
COPY --from=builder /app/proj-backend ./proj-backend
COPY --from=builder /app/internal/migrations ./internal/migrations
EXPOSE 8080
CMD ["./proj-backend"]