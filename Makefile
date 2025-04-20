.PHONY: run build setup download-tailwind dev dev-all tailwind-watch templ-watch templ download-templ serve-static

# @INFO production build and run
run: templ tailwind bundle
	@go build -tags prod -o ./bin/main . && \
		./bin/main

# @INFO development build and run
air:
	@air

dev: templ tailwind
	@echo "Starting development server with all watchers..."
	@$(MAKE) -j5 air templ-watch tailwind-watch bundle-watch browser-sync
# ----------------------------------------------------

# @INFO setup the project
setup: download-tailwind download-templ tailwind templ env log
	@go mod tidy

# ----------------------------------------------------
clean:
	@rm -rf ./bin && \
		rm -rf ./tmp && \
		rm -rf ./logs && \
		echo "Cleaned up the project."

# ----------------------------------------------------
log:
	@mkdir -p ./logs && \
		touch ./logs/app.log && \
		echo "Log file created."

env:
	@touch .env && \
		cat .env.example > .env && \
		echo "Environment file created."

# ----------------------------------------------------
bundle:
	@go mod tidy && \
		find ./scripts -type f -name '*.js' ! -name 'index.js' -exec echo "import './{}';" \; | sed 's|./scripts/||' > ./scripts/index.js && \
		npx esbuild ./scripts/index.js --bundle --minify --sourcemap --outfile=./public/script.js

bundle-watch:
	@go mod tidy && \
		find ./scripts -type f -name '*.js' ! -name 'index.js' -exec echo "import './{}';" \; | sed 's|./scripts/||' > ./scripts/index.js && \
		npx esbuild ./scripts/index.js --bundle --minify --outfile=./public/script.js --watch=forever

browser-sync:
	@echo "Starting BrowserSync..."
	@npx browser-sync start --proxy "localhost:9999" --files "./public/**/*.*" --files "./bin/main" --no-open


# ----------------------------------------------------
download-tailwind:
	@curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && \
		chmod +x tailwindcss-linux-x64 && \
		mkdir -p ./bin && \
		mv tailwindcss-linux-x64 ./bin/tailwindcss && \
		./bin/tailwindcss --version && \
		echo "Tailwind CSS downloaded and made executable."

download-templ:
	@curl -sLO https://github.com/a-h/templ/releases/latest/download/templ_Linux_x86_64.tar.gz && \
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
	@./bin/templ generate --watch
