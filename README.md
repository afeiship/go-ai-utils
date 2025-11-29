# go-ai-utils
> Go utilities for LLM-powered text tasks: keywords, summary, translation, and more.

## Installation
```sh
go get -u github.com/afeiship/aiutils
```

## Features

### 🎯 AI关键词提取 (AI Keywords Extraction)

基于 Claude AI 的智能关键词提取功能，采用面向对象的 Client 模式 API。

#### 推荐用法 - Client 模式
```bash
# 设置环境变量
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
    content := "人工智能是计算机科学的一个分支..."

    // 创建客户端（从环境变量获取API Key）
    client := aiutils.NewClientFromEnv()

    // 简单提取关键词
    keywords, err := client.KeywordsSimple(ctx, content)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("关键词:", keywords)

    // 使用完整选项
    result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
        Count:    5,
        Language: aiutils.LanguageEnglish,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("英文关键词: %v\n", result.Keywords)
}
```

#### 高级用法 - 链式配置
```go
// 链式配置客户端
client := aiutils.NewClientFromEnv().
    SetModel("glm-4.5-air").
    SetMaxTokens(1500)

// 不同调用方式
keywords1, _ := client.KeywordsSimple(ctx, content)
keywords2, _ := client.KeywordsWithCount(ctx, content, 3)
keywords3, _ := client.KeywordsWithLanguage(ctx, content, aiutils.LanguageEnglish)
result, _ := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
    Count:    8,
    Language: aiutils.LanguageMixed,
})
```

#### 特性
- 🎯 **面向对象设计**: Client + Client.Keywords 模式，更符合Go惯用法
- 🌍 **多语言支持**: 中文、英文、混合语言关键词提取
- 🔧 **灵活配置**: 支持链式配置、方法参数等多种方式
- 📦 **嵌入式模板**: 提示词模板内置，无需外部文件
- 🌐 **环境变量支持**: 优先使用环境变量，便于部署
- 🔗 **链式调用**: 流畅的API设计，支持方法链

## Project Structure

```
go-ai-utils/
├── keywords.go              # Client 主要实现
├── utils.go                 # 私有工具函数
├── assets/
│   └── strings.yml          # 嵌入的YAML提示词模板
├── examples/
│   └── keywords_example.go  # 使用示例
├── docs/
│   ├── ai-keywords.md       # AI关键词提取详细文档
│   └── 01-go-claude.md      # go-claude包文档
├── go.mod
├── go.sum
└── README.md
```

## Documentation

- [AI关键词提取详细文档](docs/ai-keywords.md)
- [示例代码](examples/keywords_example.go)
- [go-claude包文档](docs/01-go-claude.md)

## License

MIT