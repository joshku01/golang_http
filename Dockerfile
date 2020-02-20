# 第一層基底
FROM golang:1.12 as builder

ENV GO111MODULE=on

# Copy local code to the container image.
WORKDIR /app
COPY . .

COPY go.mod .
COPY go.sum .
RUN go mod download

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o golang-http

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/golang-http /golang-http

# Run the web service on container startup.
CMD ["/golang-http"]
