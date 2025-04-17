run: build
	@./bin/main

build:
	@templ generate && \
		go build -o ./bin/main .

dev:
	@templ generate && \
		air