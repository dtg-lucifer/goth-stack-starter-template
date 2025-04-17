run: build
	@./bin/main

build:
	@templ generate && \
		go build -o ./bin/main ./cmd/app/main.go
