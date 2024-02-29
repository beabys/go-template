FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/go-http-template

# Use another container after app is build
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/go-http-template /app/go-http-template

# Run the compiled app
CMD ["/app/go-http-template"]
