FROM golang:1.22-alpine

# Minimal packages useful for development (git for module fetching)
RUN apk add --no-cache ca-certificates git make

WORKDIR /app

# Copy all files
COPY . .

# Initialize go modules if go.mod doesn't exist
RUN if [ ! -f go.mod ]; then go mod init todo-cli; fi && \
    go mod tidy || true

# Default to an interactive shell so user can run commands inside container
CMD ["sh"]
