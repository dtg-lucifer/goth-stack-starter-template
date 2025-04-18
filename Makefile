.PHONY: run build setup download-tailwind dev tailwind-watch templ-watch

run: build
	@./bin/main

build:
	@templ generate && \
		go build -o ./bin/main .


setup: download-tailwind
	@go mod tidy

download-tailwind:
	@curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && \
		chmod +x tailwindcss-linux-x64 && \
		mkdir -p ./bin && \
		mv tailwindcss-linux-x64 ./bin/tailwindcss && \
		./bin/tailwindcss --version && \
		./bin/tailwindcss init && \
		echo "Tailwind CSS downloaded and made executable."

dev:
	@air

tailwind-watch:
	@./bin/tailwindcss -i ./views/layout/root.css -o ./public/styles.css --watch --minify

templ-watch:
	@templ generate --watch --proxy=http://localhost:9999