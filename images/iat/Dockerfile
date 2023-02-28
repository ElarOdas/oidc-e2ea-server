# syntax=docker/dockerfile:1.4
FROM --platform=$BUILDPLATFORM golang:1.19 AS builder

WORKDIR /code 

ENV CGO_ENABLED 1
ENV GOOS linux
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -a -ldflags '-linkmode external -extldflags "-static"' -o bin/iat main.go

CMD ["/code/bin/iat"]

# ! Dev Environment currently not needed
# FROM builder as dev-envs

# RUN <<EOF
# apk update
# apk add git
# EOF

# RUN <<EOF
# addgroup -S docker
# adduser -S --shell /bin/bash --ingroup docker vscode
# EOF

# # install Docker tools (cli, buildx, compose)
# COPY --from=gloursdocker/docker / /

# CMD ["go", "run", "main.go"]

FROM scratch
COPY --from=builder /code/bin/iat /usr/local/bin/iat

# Set default configuration parameters
ENV ALG="ES256"
ENV DEFAULT_TOKEN_PERIOD=3600
ENV MAX_TOKEN_PERIOD=2592000
ENV DB_SQLITE_FILE="/config/db.sqlite"

# Run the binary on start
CMD ["/usr/local/bin/iat"]