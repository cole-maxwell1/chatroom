ARG PORT=8080

FROM golang:1.22.0-alpine AS build

# Build binary from go source
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN GOOS=linux go build -v -o bin/chatroom cmd/main.go

FROM scratch
WORKDIR /app
# Copy binary from build step
COPY --from=build /app/bin/chatroom /app/chatroom

# Set startup options
EXPOSE ${PORT}
ENTRYPOINT ["/app/chatroom"]
