# Stage 1 - Builder: Import the golang container.
FROM golang:1.20-alpine as builder

# Install ssh client and git
RUN apk add --no-cache openssh-client git

# Set the work directory.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod ./
COPY go.sum ./

# This command will have access to the forwarded agent (if one is
# available)
RUN go mod download

# Copy the source code into the container.
COPY ./ ./

# Build the source code
RUN CGO_ENABLED=0 go build -o ./out/executable .


# Stage 2 - Runner.
FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder /app/out/executable executable

CMD [ "/app/executable" ]