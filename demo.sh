#!/bin/bash

# Demo script for Ollama Example Project
# This script demonstrates all the features of the project

set -e

echo "🚀 Ollama Example Project Demo"
echo "=============================="
echo ""

# Check if setup script exists and run it
if [ -f "check-setup.sh" ]; then
    echo "📋 Running setup verification..."
    ./check-setup.sh
    echo ""
    echo "Press Enter to continue with the demo..."
    read -r
else
    echo "⚠️  Setup verification script not found. Proceeding with demo..."
fi

echo ""
echo "🔨 Building the project..."
make build

echo ""
echo "📝 Demo Files Available:"
echo "========================"
echo "1. sample.txt - A comprehensive text about machine learning"
echo "2. example.go - Go source code for a user management API"
echo "3. README.md - This project's documentation"
echo ""

# Function to run summarization with a header
run_summary() {
    local file=$1
    local description=$2

    echo ""
    echo "📄 Summarizing: $file"
    echo "Description: $description"
    echo "----------------------------------------"

    if [ -f "$file" ]; then
        echo "Running: ./file-summarizer summarize $file"
        echo ""
        ./file-summarizer summarize "$file"
    else
        echo "❌ File $file not found!"
    fi

    echo ""
    echo "Press Enter to continue..."
    read -r
}

echo "🎯 Starting File Summarization Demo"
echo "===================================="

# Demo 1: Sample text file
run_summary "sample.txt" "Technical article about machine learning"

# Demo 2: Go source code
run_summary "example.go" "Go source code for user management API"

# Demo 3: README file
run_summary "README.md" "Project documentation and setup instructions"

echo ""
echo "🌐 API Demonstration"
echo "===================="
echo ""
echo "The following examples show how to use Ollama's REST API directly:"
echo ""

echo "1. Native Ollama API:"
echo "curl -X POST http://localhost:11434/api/generate \\"
echo "  -H 'Content-Type: application/json' \\"
echo "  -d '{"
echo "    \"model\": \"granite-code:8b\","
echo "    \"prompt\": \"Explain what is machine learning in simple terms\","
echo "    \"stream\": false"
echo "  }'"
echo ""

echo "2. OpenAI Compatible API:"
echo "curl -X POST http://localhost:11434/v1/chat/completions \\"
echo "  -H 'Content-Type: application/json' \\"
echo "  -H 'Authorization: Bearer ollama' \\"
echo "  -d '{"
echo "    \"model\": \"granite-code:8b\","
echo "    \"messages\": ["
echo "      {"
echo "        \"role\": \"user\","
echo "        \"content\": \"Explain recursion in programming\""
echo "      }"
echo "    ]"
echo "  }'"
echo ""

echo "🛠️  Available Make Commands:"
echo "============================="
make help

echo ""
echo "✨ Demo Complete!"
echo "================="
echo ""
echo "What you've seen:"
echo "• File summarization using granite-code:8b model"
echo "• Different file types (text, Go code, markdown)"
echo "• REST API usage examples"
echo "• Build and automation tools"
echo ""
echo "Try it yourself:"
echo "• Create your own files and summarize them"
echo "• Test the REST API endpoints"
echo "• Explore different Ollama models"
echo ""
echo "📚 For more information, see README.md"
echo "🐛 Report issues: https://github.com/your-repo/ollama-example"
echo ""
echo "Happy coding! 🎉"
