# --- Frontend Build ---
FROM node:25-slim AS frontend-builder

WORKDIR /app

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

# package.json で指定されているバージョンをインストール
RUN npm install -g pnpm@10.30.1

COPY package.json pnpm-lock.yaml* ./

RUN pnpm install --frozen-lockfile

COPY . .

RUN pnpm tailwindcss -i ./assets/input.css -o ./public/css/style.css
RUN mkdir -p public/js && cp node_modules/htmx.org/dist/htmx.min.js public/js/

# --- Backend Build ---
FROM golang:1.25-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY --from=frontend-builder /app/public ./public
COPY . .

RUN go tool templ generate && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/app ./cmd/app/main.go

# --- Runtime ---
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

COPY --from=backend-builder /app/bin/app .
COPY --from=backend-builder /app/data ./data
COPY --from=backend-builder /app/public ./public

ENV PORT=8080
EXPOSE 8080

CMD ["./app"]
