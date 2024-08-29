# Generate tailwindcss
FROM node:22-alpine AS generate-tailwindcss
WORKDIR /app
COPY . .
RUN npm install -g tailwindcss
RUN npx tailwindcss -i public/css/input.css -o public/css/output.css --minify

# Fetch
FROM golang:latest AS fetch-stage 
COPY go.mod go.sum ./
WORKDIR /app
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build
FROM golang:latest AS build-stage
COPY --from=generate-stage /app /app
WORKDIR /app
RUN go build -o main.exe cmd/app/main.go

EXPOSE 8000
ENTRYPOINT ["go", "run", "cmd/app/main.go"]