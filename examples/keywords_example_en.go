package main

import (
	"fmt"
	"log"

	"github.com/afeiship/go-ai-utils"
)

func main() {
	// Example text for keyword extraction
	content := "Artificial Intelligence is a branch of computer science that aims to understand the essence of intelligence and produce intelligent machines that can react in a similar way to human intelligence. The field includes research in robotics, speech recognition, image recognition, natural language processing, and expert systems."

	// Method 1: Create client with default options (uses environment variables)
	fmt.Println("=== Using Environment Variables (Recommended) ===")
	client1 := aiutils.NewClient(aiutils.NewClientOptions())

	result1, err := client1.Keywords(content)
	if err != nil {
		log.Printf("Error with environment variables: %v (Make sure ANTHROPIC_AUTH_TOKEN is set)", err)
	} else {
		fmt.Printf("Extracted %d keywords:\n", result1.Count)
		for i, keyword := range result1.Keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}

	fmt.Println("\n=== With Custom Options ===")
	// Method 2: Create client with custom options
	client2 := aiutils.NewClient(aiutils.ClientOptions{}.
		WithAPIKey("your-api-key-here").
		WithModel("claude-3-sonnet-20240229").
		WithMaxTokens(1000))

	// Extract keywords with custom options
	result2, err := client2.Keywords(content, &aiutils.KeywordsOptions{
		Count:    5,
		Language: aiutils.LanguageEnglish,
	})

	if err != nil {
		log.Printf("Error with custom client: %v", err)
	} else {
		fmt.Printf("Extracted %d keywords in English:\n", result2.Count)
		for i, keyword := range result1.Keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}

	fmt.Println("\n=== Using Fluent Interface ===")
	// Method 3: Using the fluent interface with method chaining
	client3 := aiutils.NewClient(aiutils.NewClientOptions()).
		SetOptions(aiutils.ClientOptions{}.
			WithModel("claude-3-haiku-20240307").
			WithMaxTokens(500))

	result3, err := client3.Keywords(content, &aiutils.KeywordsOptions{
		Count:    3,
		Language: aiutils.LanguageMixed,
	})

	if err != nil {
		log.Printf("Error with fluent interface: %v", err)
	} else {
		fmt.Printf("Extracted %d keywords (mixed language):\n", result3.Count)
		for i, keyword := range result3.Keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}
}