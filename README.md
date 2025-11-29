# go-ai-utils
> Go utilities for LLM-powered text tasks: keywords, summary, translation, and more.

## Installation

```bash
go get github.com/afeiship/aiutils
```

## Quick Start

```bash
# Set environment variable
export ANTHROPIC_AUTH_TOKEN="your-api-key"
export ANTHROPIC_BASE_URL="https://api.anthropic.ai"
```

```go
package main

import (
      "fmt"
    "log"
    "github.com/afeiship/aiutils"
)

func main() {
    content := "人工智能是计算机科学的一个分支，它企图了解智能的实质..."

    // Create client and extract keywords
    client := aiutils.NewClient(aiutils.NewClientOptions())
    result, err := client.Keywords(content)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("关键词:", result.Keywords)
}
```

## Advanced Usage

```go
// With options
result, err := client.Keywords(content, &aiutils.KeywordsOptions{
    Count:    5,
    Language: aiutils.LanguageEnglish,
})

// Configure client with chainable helpers
client := aiutils.NewClient(aiutils.NewClientOptions().
    WithAPIKey("api-key").
    WithModel("glm-4.5-air").
    WithMaxTokens(1500)

// Or directly
client := aiutils.NewClient(aiutils.ClientOptions{
    APIKey:    "api-key",
    Model:     "glm-4.5-air",
    MaxTokens: 1500,
})
```