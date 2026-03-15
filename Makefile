.DEFAULT_GOAL := help

.PHONY: help
help:
	@printf "%-30s %-60s\n" "[Sub command]" "[Description]"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %-60s\n", $$1, $$2}'

.PHONY: setup
setup: ## Set up
	@go mod tidy
	@pnpm install --frozen-lockfile
	@mkdir -p public/js
	@mkdir -p public/css
	@cp node_modules/htmx.org/dist/htmx.min.js public/js/
	@cp node_modules/chart.js/dist/chart.umd.js public/js/

.PHONY: air
air: ## Start the development server with live reload
	@air -c .air.toml

.PHONY: build
build: ## Build the project
	@echo "Building the project..."
	@templ generate
	@pnpm exec tailwindcss -i ./assets/input.css -o ./public/css/style.css
	@go build -o ./bin/app ./cmd/app
