# Ollama Example Project

This project demonstrates how to use Ollama for AI inference with a simple Go CLI application that can summarize files using the granite-code model.

## Quick Start

```bash
# 1. Install Ollama (macOS)
brew install ollama

# 2. Start Ollama service (keep this running)
ollama serve

# 3. Download the model (in another terminal)
ollama pull granite-code:8b

# 4. Setup and run the project
cd ollama-example
make setup
make demo
```

That's it! The demo will guide you through all the features.

## Prerequisites

- macOS (instructions provided for macOS)
- Go 1.19 or later
- Terminal access

## Installing Ollama on macOS

### Method 1: Using Homebrew (Recommended)

```bash
brew install ollama
```

### Method 2: Direct Download

1. Visit the [Ollama website](https://ollama.com)
2. Download the macOS installer
3. Run the installer and follow the prompts

### Method 3: Using curl

```bash
curl -fsSL https://ollama.com/install.sh | sh
```

## Starting Ollama Service

After installation, start the Ollama service:

```bash
ollama serve
```

This will start the Ollama server on `http://localhost:11434`. Keep this terminal window open or run it in the background.

Alternatively, if Ollama was installed using macOS installer, just lanch Ollama app itself.

## Downloading a Model

To download the granite-code:8b model, run:

```bash
ollama pull granite-code:8b
```

This will download the model to your local machine. The download may take some time depending on your internet connection.

To verify the model is installed:

```bash
ollama list
```

## Using Ollama REST API

### Method 1: Default Ollama HTTP REST Endpoint

You can interact with Ollama using its native REST API:

```bash
curl -X POST http://localhost:11434/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "model": "granite-code:8b",
    "prompt": "Explain what is machine learning in simple terms",
    "stream": false
  }'
```

For chat-style conversations:

```bash
curl -X POST http://localhost:11434/api/chat \
  -H "Content-Type: application/json" \
  -d '{
    "model": "granite-code:8b",
    "messages": [
      {
        "role": "user",
        "content": "What is the difference between arrays and slices in Go?"
      }
    ],
    "stream": false
  }'
```

### Method 2: OpenAI Compatible Endpoint

Ollama provides an OpenAI-compatible API endpoint. You can use it with OpenAI client libraries:

```bash
curl -X POST http://localhost:11434/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ollama" \
  -d '{
    "model": "granite-code:8b",
    "messages": [
      {
        "role": "user",
        "content": "Explain recursion in programming"
      }
    ]
  }'
```

You can also use this endpoint with any OpenAI-compatible client by setting:
- Base URL: `http://localhost:11434/v1`
- API Key: `ollama` (or any string, it's not validated)

## Go CLI Application

This project includes a simple Go CLI application that uses the granite-code model to summarize files.

### Setup Verification

Before getting started, you can run the setup verification script to check if everything is properly configured:

```bash
./check-setup.sh
```

This script will check:
- Go installation
- Ollama installation and service status
- granite-code:8b model availability
- Project dependencies and build status
- Test files presence

### Demo

To see all features in action, run the interactive demo:

```bash
./demo.sh
```

This demo will:
- Verify your setup
- Build the project
- Demonstrate file summarization with different file types
- Show REST API usage examples
- Display available commands

### Quick Setup

For a quick setup, you can use the provided Makefile:

```bash
# Setup everything (install dependencies and download model)
make setup

# Start Ollama service (in a separate terminal)
ollama serve

# Build and run with the sample file
make run
```

### Manual Setup

#### Building the Application

```bash
cd ollama-example
go mod tidy
go build -o file-summarizer
```

#### Running the Application

To summarize a file:

```bash
./file-summarizer summarize path/to/your/file.txt
```

Example:

```bash
./file-summarizer summarize README.md
```

### Available Make Commands

```bash
make build       # Build the CLI application
make clean       # Clean build artifacts
make deps        # Install Go dependencies
make run         # Build and run with sample.txt
make run-readme  # Build and run with README.md
make run-example # Build and run with example.go
make test        # Run tests
make check-ollama # Check if Ollama service is running
make pull-model  # Download the granite-code:8b model
make list-models # List installed Ollama models
make setup       # Setup dependencies and model
make help        # Show help message
```

### Usage Examples

```bash
# Summarize a Go source file
./file-summarizer summarize main.go

# Summarize a text document
./file-summarizer summarize document.txt

# Summarize this README
./file-summarizer summarize README.md

# Summarize the provided sample file
./file-summarizer summarize sample.txt

# Summarize the example Go code
./file-summarizer summarize example.go
```

## Project Structure

```
ollama-example/
├── README.md           # This file
├── main.go            # Main CLI application
├── example.go         # Example Go code for testing
├── go.mod             # Go module file
├── go.sum             # Go dependencies
├── Makefile           # Build and run commands
├── .gitignore         # Git ignore file
├── sample.txt         # Sample file for testing
├── check-setup.sh     # Setup verification script
├── demo.sh            # Interactive demo script
└── file-summarizer    # Built executable (after running make build)
```

## Dependencies

This project uses:
- [langchaingo](https://github.com/tmc/langchaingo) - Go library for building applications with LLMs
- [cobra](https://github.com/spf13/cobra) - CLI framework for Go

## Troubleshooting

### Ollama Service Not Running

If you get connection errors, make sure Ollama is running:

```bash
ollama serve
```

### Model Not Found

If you get a "model not found" error, pull the model:

```bash
ollama pull granite-code:8b
```

### Port Conflicts

If port 11434 is busy, you can start Ollama on a different port:

```bash
OLLAMA_HOST=0.0.0.0:11435 ollama serve
```

Then update the code to use the new port.

## Additional Resources

- [Ollama Documentation](https://github.com/ollama/ollama)
- [LangChain Go Documentation](https://github.com/tmc/langchaingo)
- [Granite Code Model Information](https://ollama.com/library/granite-code)
