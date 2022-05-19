FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY go.mod go.sum ./

RUN go mod download

# Build the binary.
COPY . .


RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -o /go/bin/spotify-server .

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/spotify-server /spotify-server
# Run the hello binary.
ENTRYPOINT ["/spotify-server"]