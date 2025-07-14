.PHONY: build clean run test help

# Build the CLI application
build:
	go build -o file-summarizer

# Clean build artifacts
clean:
	rm -f file-summarizer

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run the application with sample file
run: build
	./file-summarizer summarize sample.txt

# Run the application with README
run-readme: build
	./file-summarizer summarize README.md

# Run the application with example Go code
run-example: build
	./file-summarizer summarize example.go

# Run the interactive demo
demo: build
	./demo.sh

# Test the application
test:
	go test ./...

# Check if Ollama is running
check-ollama:
	@echo "Checking if Ollama is running..."
	@curl -s http://localhost:11434/api/tags > /dev/null && echo "✓ Ollama is running" || echo "✗ Ollama is not running. Please start it with 'ollama serve'"

# Pull the granite-code model
pull-model:
	ollama pull granite-code:8b

# List available models
list-models:
	ollama list

# Setup everything needed to run the project
setup: deps pull-model
	@echo "Setup complete! Make sure to run 'ollama serve' in another terminal."

# Show help
help:
	@echo "Available commands:"
	@echo "  build       - Build the CLI application"
	@echo "  clean       - Clean build artifacts"
	@echo "  deps        - Install Go dependencies"
	@echo "  run         - Build and run with sample.txt"
	@echo "  run-readme  - Build and run with README.md"
	@echo "  run-example - Build and run with example.go"
	@echo "  demo        - Run interactive demo"
	@echo "  test        - Run tests"
	@echo "  check-ollama - Check if Ollama service is running"
	@echo "  pull-model  - Download the granite-code:8b model"
	@echo "  list-models - List installed Ollama models"
	@echo "  setup       - Setup dependencies and model"
	@echo "  help        - Show this help message"
