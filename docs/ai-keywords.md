# AI关键词提取工具

基于 go-claude 包实现的 AI 关键词提取功能，采用面向对象的 Client 模式设计。

## 功能特性

- 🤖 使用 Claude AI 智能提取关键词
- 🌍 支持中文、英文和混合语言
- ⚙️ 可配置关键词数量
- 🎯 面向对象的 Client 设计
- 🔧 支持链式配置和方法调用
- 🌐 环境变量和直接传入参数支持
- 📦 嵌入式提示词模板，无需外部文件

## 安装

```bash
go get github.com/afeiship/go-ai-utils
```

## 使用示例

### 基础用法 - Client 模式

#### 方法1: 使用环境变量创建Client (推荐)

```bash
# 设置环境变量
export ANTHROPIC_AUTH_TOKEN="your-api-key"
export ANTHROPIC_BASE_URL="https://api.anthropic.com"  # 可选
```

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/afeiship/go-aiutils"
)

func main() {
    ctx := context.Background()
    content := "人工智能是计算机科学的一个分支..."

    // 从环境变量创建客户端
    client := aiutils.NewClientFromEnv()

    // 简单提取关键词
    keywords, err := client.KeywordsSimple(ctx, content)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("提取的关键词:", keywords)
}
```

#### 方法2: 直接传入API Key

```go
apiKey := "your-api-key"
client := aiutils.NewClient(apiKey)

keywords, err := client.KeywordsSimple(ctx, content)
if err != nil {
    log.Fatal(err)
}
```

#### 方法3: 使用完整选项创建Client

```go
client := aiutils.NewClientWithOptions(
    "your-api-key",                    // API Key
    "https://api.anthropic.com",       // Base URL
    "glm-4.5-air",                    // Model
    1500,                             // Max Tokens
)
```

### 高级用法

#### 链式配置Client

```go
// 链式配置，支持方法调用
client := aiutils.NewClientFromEnv().
    SetModel("glm-4.5-air").
    SetMaxTokens(1500).
    SetBaseURL("https://custom-api.com")
```

#### 不同的关键词提取方法

```go
// 1. 简单提取（使用默认配置）
keywords, err := client.KeywordsSimple(ctx, content)

// 2. 指定关键词数量
keywords, err := client.KeywordsWithCount(ctx, content, 3)

// 3. 指定语言
keywords, err := client.KeywordsWithLanguage(ctx, content, aiutils.LanguageEnglish)

// 4. 使用完整选项（推荐）
result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
    Count:    8,
    Language: aiutils.LanguageMixed,
})

fmt.Printf("提取到 %d 个关键词\n", result.Count)
for i, keyword := range result.Keywords {
    fmt.Printf("%d. %s\n", i+1, keyword)
}
```

#### 完整示例

```go
func main() {
    ctx := context.Background()
    content := "人工智能是计算机科学的一个分支..."

    // 创建并配置客户端
    client := aiutils.NewClientFromEnv().
        SetModel("glm-4.5-air").
        SetMaxTokens(1500)

    // 多种提取方式
    methods := []struct {
        name string
        fn   func() ([]string, error)
    }{
        {
            "默认提取",
            func() ([]string, error) {
                return client.KeywordsSimple(ctx, content)
            },
        },
        {
            "提取3个关键词",
            func() ([]string, error) {
                return client.KeywordsWithCount(ctx, content, 3)
            },
        },
        {
            "提取英文关键词",
            func() ([]string, error) {
                return client.KeywordsWithLanguage(ctx, content, aiutils.LanguageEnglish)
            },
        },
        {
            "提取8个混合语言关键词",
            func() ([]string, error) {
                result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
                    Count:    8,
                    Language: aiutils.LanguageMixed,
                })
                return result.Keywords, err
            },
        },
    }

    for _, method := range methods {
        keywords, err := method.fn()
        if err != nil {
            log.Printf("%s 错误: %v", method.name, err)
            continue
        }

        fmt.Printf("=== %s ===\n", method.name)
        for i, keyword := range keywords {
            fmt.Printf("%d. %s\n", i+1, keyword)
        }
        fmt.Println()
    }
}
```

## API 参考

### Client 构造函数

#### NewClient(apiKey string) *Client
**基础构造函数** - 创建新的AI客户端

```go
client := aiutils.NewClient("your-api-key")
```

#### NewClientFromEnv() *Client
**环境变量构造函数** - 从环境变量创建AI客户端

```go
client := aiutils.NewClientFromEnv() // 使用 ANTHROPIC_AUTH_TOKEN
```

#### NewClientWithOptions(apiKey, baseURL, model string, maxTokens int) *Client
**完整选项构造函数** - 使用完整选项创建AI客户端

```go
client := aiutils.NewClientWithOptions(
    "your-api-key",                    // API Key
    "https://api.anthropic.com",       // Base URL
    "glm-4.5-air",                    // Model
    1500,                             // Max Tokens
)
```

### Client 链式方法

#### SetBaseURL(baseURL string) *Client
设置API基础URL

```go
client := client.SetBaseURL("https://custom-api.com")
```

#### SetModel(model string) *Client
设置使用的模型

```go
client := client.SetModel("glm-4.5-air")
```

#### SetMaxTokens(maxTokens int) *Client
设置最大token数

```go
client := client.SetMaxTokens(1500)
```

### Client 关键词提取方法

#### Keywords(ctx, content, options) (*KeywordsResult, error)
**完整方法** - 使用选项进行关键词提取

参数：
- `ctx context.Context`: 上下文
- `content string`: 要分析的文本内容
- `options *KeywordsOptions`: 可选配置对象

返回：
- `*KeywordsResult`: 包含关键词列表、数量和语言的结果
- `error`: 错误信息

```go
result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
    Count:    8,
    Language: aiutils.LanguageEnglish,
})
```

#### KeywordsSimple(ctx, content) ([]string, error)
**简化方法** - 使用默认配置进行关键词提取

```go
keywords, err := client.KeywordsSimple(ctx, content)
```

#### KeywordsWithCount(ctx, content, count) ([]string, error)
**指定数量方法** - 提取指定数量的关键词

```go
keywords, err := client.KeywordsWithCount(ctx, content, 3)
```

#### KeywordsWithLanguage(ctx, content, language) ([]string, error)
**指定语言方法** - 提取指定语言的关键词

```go
keywords, err := client.KeywordsWithLanguage(ctx, content, aiutils.LanguageEnglish)
```

### KeywordsOptions 配置对象

```go
type KeywordsOptions struct {
    // Count 返回关键词的数量，默认5
    Count int
    // Language 返回关键词的语言，默认中文
    Language Language
}
```

### KeywordsResult 结果对象

```go
type KeywordsResult struct {
    // Keywords 提取的关键词列表
    Keywords []string
    // Count 关键词数量
   Count int
    // Language 结果语言
    Language Language
}
```

### Language 常量

- `LanguageChinese`: 中文关键词
- `LanguageEnglish`: 英文关键词
- `LanguageMixed`: 混合语言关键词

### 环境变量

- `ANTHROPIC_AUTH_TOKEN`: Claude API密钥
- `ANTHROPIC_BASE_URL`: API基础URL (可选)

### 配置优先级

1. **API Key**: 构造函数参数 > 环境变量 ANTHROPIC_AUTH_TOKEN
2. **Base URL**: SetBaseURL() > 构造函数参数 > 环境变量 > 默认值
3. **Model**: SetModel() > 构造函数参数 > 默认值
4. **MaxTokens**: SetMaxTokens() > 构造函数参数 > 默认值
5. **Count**: KeywordsOptions.Count > 默认值（从 assets/strings.yml）
6. **Language**: KeywordsOptions.Language > 默认值（从 assets/strings.yml）

### 提示词模板

提示词模板存储在 `assets/strings.yml` 文件中，使用 Go embed 方式嵌入到二进制文件中：

- 中文关键词提取模板
- 英文关键词提取模板
- 混合语言关键词提取模板
- 默认配置参数
- 语言指令映射

#### 模板变量

提示词模板支持以下变量：
- `{{.Count}}`: 要提取的关键词数量
- `{{.Content}}`: 要分析的文本内容

#### 文件结构

```
aiutils/
├── keywords.go              # Client 主要实现
├── utils.go                 # 私有工具函数
└── assets/
    └── strings.yml          # 嵌入的YAML提示词模板配置
```

## 注意事项

1. 需要有效的 Claude API 密钥
2. 确保网络连接正常
3. 内容文本不能为空
4. API 调用需要消耗 tokens