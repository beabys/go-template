FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-http-template ./cmd/server/main.go

# Use another container after app is build
FROM alpine:latest

RUN addgroup -g 10000 appgroup
RUN adduser \
    --disabled-password \
    --gecos "" \
    --ingroup appgroup \
    --no-create-home \
    --uid 10000 \
    appuser

COPY --from=builder /app/go-http-template .

RUN mkdir -p /etc/config

RUN chown -R appuser:appgroup /etc/config

USER 10000
EXPOSE 3001

# Run the compiled app
CMD ["/go-http-template"]
