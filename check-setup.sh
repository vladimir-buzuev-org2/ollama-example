#!/bin/bash

# Setup verification script for Ollama Example Project
# This script checks if all prerequisites are met to run the project

set -e

echo "🔍 Checking Ollama Example Project Setup..."
echo "================================================="

# Check if Go is installed
echo -n "Checking Go installation... "
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}')
    echo "✅ $GO_VERSION"
else
    echo "❌ Go not found. Please install Go 1.19 or later."
    exit 1
fi

# Check if Ollama is installed
echo -n "Checking Ollama installation... "
if command -v ollama &> /dev/null; then
    echo "✅ Ollama installed"
else
    echo "❌ Ollama not found. Please install Ollama first."
    echo "   Run: brew install ollama"
    exit 1
fi

# Check if Ollama service is running
echo -n "Checking Ollama service... "
if curl -s http://localhost:11434/api/tags > /dev/null 2>&1; then
    echo "✅ Ollama service is running"
else
    echo "⚠️  Ollama service not running"
    echo "   Please start it with: ollama serve"
fi

# Check if granite-code model is available
echo -n "Checking granite-code:8b model... "
if ollama list | grep -q "granite-code:8b"; then
    echo "✅ granite-code:8b model available"
else
    echo "⚠️  granite-code:8b model not found"
    echo "   Download it with: ollama pull granite-code:8b"
fi

# Check if project dependencies are installed
echo -n "Checking Go dependencies... "
if [ -f "go.sum" ]; then
    echo "✅ Dependencies installed"
else
    echo "⚠️  Dependencies not installed"
    echo "   Run: go mod tidy"
fi

# Check if binary is built
echo -n "Checking compiled binary... "
if [ -f "file-summarizer" ]; then
    echo "✅ Binary built"
else
    echo "⚠️  Binary not found"
    echo "   Build it with: make build"
fi

# Check test files
echo -n "Checking test files... "
MISSING_FILES=()
for file in "sample.txt" "example.go" "README.md"; do
    if [ ! -f "$file" ]; then
        MISSING_FILES+=("$file")
    fi
done

if [ ${#MISSING_FILES[@]} -eq 0 ]; then
    echo "✅ All test files present"
else
    echo "⚠️  Missing files: ${MISSING_FILES[*]}"
fi

echo ""
echo "📋 Setup Summary:"
echo "=================="

# Overall status
WARNINGS=0
ERRORS=0

# Check critical components
if ! command -v go &> /dev/null; then
    echo "❌ Go is required but not installed"
    ((ERRORS++))
fi

if ! command -v ollama &> /dev/null; then
    echo "❌ Ollama is required but not installed"
    ((ERRORS++))
fi

if ! curl -s http://localhost:11434/api/tags > /dev/null 2>&1; then
    echo "⚠️  Ollama service is not running"
    ((WARNINGS++))
fi

if ! ollama list | grep -q "granite-code:8b"; then
    echo "⚠️  granite-code:8b model not downloaded"
    ((WARNINGS++))
fi

if [ ! -f "file-summarizer" ]; then
    echo "⚠️  Project not built"
    ((WARNINGS++))
fi

if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
    echo "🎉 Everything looks good! You're ready to use the project."
    echo ""
    echo "Quick start:"
    echo "  ./file-summarizer summarize sample.txt"
elif [ $ERRORS -eq 0 ]; then
    echo "⚠️  Setup mostly complete, but some optional steps remain."
    echo ""
    echo "To complete setup:"
    if ! curl -s http://localhost:11434/api/tags > /dev/null 2>&1; then
        echo "  1. Start Ollama: ollama serve"
    fi
    if ! ollama list | grep -q "granite-code:8b"; then
        echo "  2. Download model: ollama pull granite-code:8b"
    fi
    if [ ! -f "file-summarizer" ]; then
        echo "  3. Build project: make build"
    fi
else
    echo "❌ Critical components missing. Please install missing requirements."
fi

echo ""
echo "For detailed setup instructions, see README.md"
