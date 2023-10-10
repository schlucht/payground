install:
	@docker-compose up -d

build:
	@echo "Building app..."
	@go build -o dist/todo ./cmd
	@echo "App built!"

run: build
	@./dist/todo

