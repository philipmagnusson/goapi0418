# Use multistage build.
FROM golang:alpine AS builder
WORKDIR /src

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . .

# Build the application.
RUN go build -o /cmd/site

FROM scratch
COPY --from=builder /cmd/site /site
COPY config.yml .

EXPOSE 8080
ENTRYPOINT [ "/site" ]