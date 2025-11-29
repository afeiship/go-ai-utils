package aiutils

import (
	"context"
	"fmt"
)

// Language 支持的语言类型
type Language string

const (
	LanguageChinese Language = "chinese"
	LanguageEnglish Language = "english"
	LanguageMixed   Language = "mixed"
)

// Client AI客户端
type Client struct {
	apiKey    string
	baseURL   string
	model     string
	maxTokens int
}

// KeywordsOptions 关键词提取选项
type KeywordsOptions struct {
	// Count 返回关键词的数量，默认5
	Count int
	// Language 返回关键词的语言，默认中文
	Language Language
}

// KeywordsResult 关键词提取结果
type KeywordsResult struct {
	// Keywords 提取的关键词列表
	Keywords []string
	// Count 关键词数量
	Count int
	// Language 结果语言
	Language Language
}

// NewClient 创建新的AI客户端
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:    apiKey,
		baseURL:   getStringConfig().Defaults.BaseURL,
		model:     getStringConfig().Defaults.Model,
		maxTokens: getStringConfig().Defaults.MaxTokens,
	}
}

// NewClientWithOptions 使用选项创建AI客户端
func NewClientWithOptions(apiKey, baseURL, model string, maxTokens int) *Client {
	client := NewClient(apiKey)
	if baseURL != "" {
		client.baseURL = baseURL
	}
	if model != "" {
		client.model = model
	}
	if maxTokens > 0 {
		client.maxTokens = maxTokens
	}
	return client
}

// NewClientFromEnv 从环境变量创建AI客户端
func NewClientFromEnv() *Client {
	return NewClient("")
}

// SetBaseURL 设置API基础URL
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

// SetModel 设置模型
func (c *Client) SetModel(model string) *Client {
	c.model = model
	return c
}

// SetMaxTokens 设置最大token数
func (c *Client) SetMaxTokens(maxTokens int) *Client {
	c.maxTokens = maxTokens
	return c
}

// Keywords 从文本中提取关键词
func (c *Client) Keywords(ctx context.Context, content string, options *KeywordsOptions) (*KeywordsResult, error) {
	if content == "" {
		return nil, fmt.Errorf("content cannot be empty")
	}

	// 处理选项
	if options == nil {
		options = &KeywordsOptions{}
	}

	// 设置默认值
	if options.Count == 0 {
		options.Count = getStringConfig().Defaults.Count
	}
	if options.Language == "" {
		options.Language = Language(getStringConfig().Defaults.Language)
	}

	// 构建提示词
	prompt, err := buildPrompt(content, options.Language, options.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to build prompt: %w", err)
	}

	// 创建Claude客户端
	claudeClient, err := createClaudeClient(c.apiKey, c.baseURL, c.model, c.maxTokens)
	if err != nil {
		return nil, fmt.Errorf("failed to create Claude client: %w", err)
	}

	// 调用Claude API
	response, err := claudeClient.SimplePrompt(prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to call Claude API: %w", err)
	}

	// 解析响应
	keywords, err := parseKeywordsResponse(response, options.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &KeywordsResult{
		Keywords: keywords,
		Count:    len(keywords),
		Language: options.Language,
	}, nil
}

// KeywordsSimple 简化版关键词提取
func (c *Client) KeywordsSimple(ctx context.Context, content string) ([]string, error) {
	result, err := c.Keywords(ctx, content, nil)
	if err != nil {
		return nil, err
	}
	return result.Keywords, nil
}

// KeywordsWithCount 指定关键词数量的提取
func (c *Client) KeywordsWithCount(ctx context.Context, content string, count int) ([]string, error) {
	result, err := c.Keywords(ctx, content, &KeywordsOptions{
		Count: count,
	})
	if err != nil {
		return nil, err
	}
	return result.Keywords, nil
}

// KeywordsWithLanguage 指定语言的关键词提取
func (c *Client) KeywordsWithLanguage(ctx context.Context, content string, language Language) ([]string, error) {
	result, err := c.Keywords(ctx, content, &KeywordsOptions{
		Language: language,
	})
	if err != nil {
		return nil, err
	}
	return result.Keywords, nil
}