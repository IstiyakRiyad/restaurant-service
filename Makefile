build:
	@go build -o bin/restaurant

run: build
	@./bin/restaurant

dev: 
	@air start

test:
	@go test -v ./...



