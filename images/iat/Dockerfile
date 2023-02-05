#### Build ####

# Generate build container
FROM golang:1.19 AS build

# Create working directory for source files
WORKDIR /app

# Copy source files into container
COPY go ./go
COPY *.go .
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod tidy

# Compile application to single binary file 'iat'
RUN CGO_ENABLED=1 go build -a -installsuffix cgo -o iat

#### Deploy ####

# Generate runtime container
FROM scratch

# Create working directory for binary
WORKDIR /

# Copy compiled binary from build container
COPY --from=build /app/iat iat

# Set default configuration parameters
ENV ALG="ES256"
ENV DEFAULT_TOKEN_PERIOD=3600
ENV MAX_TOKEN_PERIOD=2592000
ENV DB_SQLITE_FILE="/config/db.sqlite"

# Run the binary on start
CMD ["./iat"]
