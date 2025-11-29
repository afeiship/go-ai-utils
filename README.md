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

    // 使用完整选项
    result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
        Count:    5,
        Language: aiutils.LanguageEnglish,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("提取到 %d 个英文关键词: %v\n", result.Count, result.Keywords)
}
```

#### 高级用法 - 配置选项
```go
// 创建时传入选项
client := aiutils.NewClient("your-api-key", aiutils.ClientOptions{
    Model:     "glm-4.5-air",
    MaxTokens: 1500,
})

// 或者使用SetOptions方法
client = client.SetOptions(aiutils.ClientOptions{
    BaseURL:   "https://custom-api.com",
    Model:     "glm-4.5-air",
    MaxTokens: 2000,
})

// 完整的Keywords方法调用
result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
    Count:    8,
    Language: aiutils.LanguageMixed,
})

// 使用默认配置
result, err := client.Keywords(ctx, content, nil)

// 指定数量
result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
    Count: 3,
})

// 指定语言
result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
    Language: aiutils.LanguageEnglish,
})
```

#### 特性
- 🎯 **面向对象设计**: Client + Client.Keywords 模式，更符合Go惯用法
- 🌍 **多语言支持**: 中文、英文、混合语言关键词提取
- 🔧 **灵活配置**: 支持ClientOptions构造函数参数和SetOptions方法
- 📦 **嵌入式模板**: YAML提示词模板内置，无需外部文件
- 🌐 **环境变量支持**: 优先使用环境变量，便于部署
- 🎛️ **统一接口**: 单一的Keywords方法，简化API设计

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