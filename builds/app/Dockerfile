FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/applcation

# Use another container after app is build
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/application /app/application

# Run the compiled app
CMD ["/app/application"]
