# 第一層基底
FROM golang:1.14 as builder

ENV GO111MODULE=on
ENV GOPROXY http://127.0.0.1:3000

WORKDIR /app/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

# Copy local code to the container image.
WORKDIR /app/release

ADD . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -installsuffix cgo -o golang-http

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates git gcc g++

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/release/golang-http /golang-http

# Run the web service on container startup.
CMD ["/golang-http"]
