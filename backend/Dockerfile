# Build stage
FROM golang:1.16.2-alpine AS build
ENV CGO_ENABLED=0
WORKDIR /go/src/app
RUN apk add git && go get github.com/swaggo/swag/cmd/swag
COPY go.* ./
RUN go mod download
COPY . ./
# Generate swagger docs
RUN go generate ./...
# Enable parallel builds and strip all debug symbols from resulting binaries
RUN mkdir /build && go build -gcflags="-c 4" -ldflags="-s -w" -o /build ./...

# Distroless for minimal footprint
FROM gcr.io/distroless/static-debian10 AS production
COPY --from=build /build /
