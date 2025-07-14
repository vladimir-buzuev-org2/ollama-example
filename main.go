package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

var rootCmd = &cobra.Command{
	Use:   "file-summarizer",
	Short: "A CLI tool to summarize files using Ollama and granite-code model",
	Long: `A simple CLI application that uses Ollama with the granite-code model
to summarize the contents of files. The summary is output to stdout.`,
}

var summarizeCmd = &cobra.Command{
	Use:   "summarize [file]",
	Short: "Summarize a file",
	Long: `Summarize the contents of a file using the granite-code model.
The file content will be read and sent to the model for summarization.`,
	Args: cobra.ExactArgs(1),
	Run:  summarizeFile,
}

func init() {
	rootCmd.AddCommand(summarizeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func summarizeFile(cmd *cobra.Command, args []string) {
	filePath := args[0]

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filePath)
	}

	// Read file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Get file extension for context
	ext := filepath.Ext(filePath)
	fileName := filepath.Base(filePath)

	// Initialize Ollama client
	llm, err := ollama.New(
		ollama.WithModel("granite-code:8b"),
		ollama.WithServerURL("http://localhost:11434"),
	)
	if err != nil {
		log.Fatalf("Error initializing Ollama client: %v", err)
	}

	// Create the prompt
	prompt := createSummarizationPrompt(string(content), fileName, ext)

	fmt.Printf("Summarizing file: %s\n", filePath)
	fmt.Println("Please wait...\n")

	// Generate summary
	ctx := context.Background()
	response, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatalf("Error generating summary: %v", err)
	}

	// Output the summary
	fmt.Println("=== SUMMARY ===")
	fmt.Println(response)
}

func createSummarizationPrompt(content, fileName, fileExtension string) string {
	var fileTypeContext string

	switch fileExtension {
	case ".go":
		fileTypeContext = "This is a Go source code file."
	case ".py":
		fileTypeContext = "This is a Python source code file."
	case ".js", ".ts":
		fileTypeContext = "This is a JavaScript/TypeScript source code file."
	case ".java":
		fileTypeContext = "This is a Java source code file."
	case ".cpp", ".cc", ".cxx":
		fileTypeContext = "This is a C++ source code file."
	case ".c":
		fileTypeContext = "This is a C source code file."
	case ".rs":
		fileTypeContext = "This is a Rust source code file."
	case ".md":
		fileTypeContext = "This is a Markdown documentation file."
	case ".txt":
		fileTypeContext = "This is a plain text file."
	case ".json":
		fileTypeContext = "This is a JSON data file."
	case ".yaml", ".yml":
		fileTypeContext = "This is a YAML configuration file."
	case ".xml":
		fileTypeContext = "This is an XML file."
	case ".html":
		fileTypeContext = "This is an HTML file."
	case ".css":
		fileTypeContext = "This is a CSS stylesheet file."
	case ".sql":
		fileTypeContext = "This is a SQL database script file."
	default:
		fileTypeContext = "This is a text file."
	}

	prompt := fmt.Sprintf(`Please provide a clear and concise summary of the following file:

File: %s
%s

Content:
%s

Please summarize:
1. What this file is about
2. Key components, functions, or sections (if applicable)
3. Main purpose or functionality
4. Any important details or notable features

Keep the summary informative but concise.`, fileName, fileTypeContext, content)

	return prompt
}
