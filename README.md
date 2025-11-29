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
```

```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/afeiship/aiutils"
)

func main() {
    ctx := context.Background()
    content := "人工智能是计算机科学的一个分支，它企图了解智能的实质..."

    // Create client and extract keywords
    client := aiutils.NewClientFromEnv()
    result, err := client.Keywords(ctx, content)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("关键词:", result.Keywords)
}
```

## Advanced Usage

```go
// With options
result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
    Count:    5,
    Language: aiutils.LanguageEnglish,
})

// Configure client
client := aiutils.NewClient("api-key", aiutils.ClientOptions{
    Model:     "glm-4.5-air",
    MaxTokens: 1500,
})
```