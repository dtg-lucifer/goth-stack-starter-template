.PHONY: run build setup download-tailwind dev tailwind-watch templ-watch templ download-templ

# @INFO production build and run
run: build
	@./bin/main

build: tailwind templ
	@go build -o ./bin/main .

# @INFO development build and run
dev:
	@air
# ----------------------------------------------------

# @INFO setup the project
setup: download-tailwind download-templ tailwind templ env
	@go mod tidy

env:
	@touch .env && \
		cat .env.example > .env && \
		echo "Environment file created."

download-tailwind:
	@curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && \
		chmod +x tailwindcss-linux-x64 && \
		mkdir -p ./bin && \
		mv tailwindcss-linux-x64 ./bin/tailwindcss && \
		./bin/tailwindcss --version && \
		echo "Tailwind CSS downloaded and made executable."

download-templ:
	@curl -sLO https://github.com/a-h/templ/releases/download/latest/templ_Linux_x86_64.tar.gz && \
		mkdir -p tmp && mkdir -p bin && \
		tar -xzf templ_Linux_x86_64.tar.gz -C ./tmp && \
		mv ./tmp/templ ./bin/templ && \
		chmod +x ./bin/templ && \
		rm -rf tmp && \
		rm templ_Linux_x86_64.tar.gz && \
		./bin/templ --version && \
		echo "Templ downloaded and extracted."

tailwind:
	@./bin/tailwindcss -i ./views/layout/root.css -o ./public/styles.css --minify && \
		echo "Tailwind CSS files generated."

templ:
	@./bin/templ generate && \
		echo "Templ files generated."

tailwind-watch:
	@./bin/tailwindcss -i ./views/layout/root.css -o ./public/styles.css --watch --minify

templ-watch:
	@./bin/templ generate --watch --proxy=http://localhost:9999