# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
COPY . .
RUN ["templ", "generate"]



# Generate tailwindcss
FROM node:22-alpine AS generate-tailwindcss
WORKDIR /app
RUN npm install -g tailwindcss
RUN npx tailwindcss -i public/css/input.css -o public/css/output.css --minify

FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN go build -o main.exe cmd/app/main.go

EXPOSE 8000
ENTRYPOINT ["go", "run", "cmd/app/main.go"]