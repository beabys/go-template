FROM golang:1.23.0-alpine3.20 as build
WORKDIR /helloworld
# Copy dependencies list
COPY . .
RUN go mod download
# Build the app
RUN go build -o main ./cmd/serverless/main.go
# Copy artifacts to a clean image
FROM alpine:3.16
COPY --from=build /helloworld/main /main
COPY --from=build /helloworld/builds/assets/config.yaml /config.yaml
ENTRYPOINT [ "/main" ]