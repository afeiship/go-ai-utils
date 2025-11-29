package aiutils

import (
	"fmt"
	"os"
)

// Language 支持的语言类型
type Language string

const (
	LanguageChinese Language = "chinese"
	LanguageEnglish Language = "english"
	LanguageMixed   Language = "mixed"
)

// ClientOptions 客户端配置选项
type ClientOptions struct {
	// APIKey Claude API密钥
	APIKey string
	// BaseURL API基础URL
	BaseURL string
	// Model 使用的模型
	Model string
	// MaxTokens 最大token数
	MaxTokens int
}

// Client AI客户端
type Client struct {
	options ClientOptions
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
func NewClient(options ClientOptions) *Client {
	defaults := getStringConfig().Defaults

	// APIKey: 用户提供 > 环境变量 > 默认值
	if options.APIKey == "" {
		options.APIKey = os.Getenv("ANTHROPIC_AUTH_TOKEN")
	}

	// BaseURL: 用户提供 > 环境变量 > 默认值
	if options.BaseURL == "" {
		options.BaseURL = os.Getenv("ANTHROPIC_BASE_URL")
	}
	if options.BaseURL == "" {
		options.BaseURL = defaults.BaseURL
	}

	// Model: 用户提供 > 默认值
	if options.Model == "" {
		options.Model = defaults.Model
	}

	// MaxTokens: 用户提供 > 默认值
	if options.MaxTokens == 0 {
		options.MaxTokens = defaults.MaxTokens
	}

	return &Client{
		options: options,
	}
}

// NewClientOptions 创建新的客户端选项
func NewClientOptions() ClientOptions {
	return ClientOptions{}
}

// WithAPIKey 设置API Key
func (opts ClientOptions) WithAPIKey(apiKey string) ClientOptions {
	opts.APIKey = apiKey
	return opts
}

// WithBaseURL 设置BaseURL
func (opts ClientOptions) WithBaseURL(baseURL string) ClientOptions {
	opts.BaseURL = baseURL
	return opts
}

// WithModel 设置Model
func (opts ClientOptions) WithModel(model string) ClientOptions {
	opts.Model = model
	return opts
}

// WithMaxTokens 设置MaxTokens
func (opts ClientOptions) WithMaxTokens(maxTokens int) ClientOptions {
	opts.MaxTokens = maxTokens
	return opts
}

// SetOptions 设置客户端选项
func (client *Client) SetOptions(options ClientOptions) *Client {
	if options.APIKey != "" {
		client.options.APIKey = options.APIKey
	}
	if options.BaseURL != "" {
		client.options.BaseURL = options.BaseURL
	}
	if options.Model != "" {
		client.options.Model = options.Model
	}
	if options.MaxTokens > 0 {
		client.options.MaxTokens = options.MaxTokens
	}
	return client
}

// Keywords 从文本中提取关键词
func (client *Client) Keywords(content string, options ...*KeywordsOptions) (*KeywordsResult, error) {
	if content == "" {
		return nil, fmt.Errorf("content cannot be empty")
	}

	// 处理选项
	var opts *KeywordsOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = &KeywordsOptions{}
	}

	// 设置默认值
	if opts.Count == 0 {
		opts.Count = getStringConfig().Defaults.Count
	}
	if opts.Language == "" {
		opts.Language = Language(getStringConfig().Defaults.Language)
	}

	// 构建提示词
	prompt, err := buildPrompt(content, opts.Language, opts.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to build prompt: %w", err)
	}

	// 创建Claude客户端
	claudeClient, err := createClaudeClient(client.options.APIKey, client.options.BaseURL, client.options.Model, client.options.MaxTokens)
	if err != nil {
		return nil, fmt.Errorf("failed to create Claude client: %w", err)
	}

	// 调用Claude API
	response, err := claudeClient.SimplePrompt(prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to call Claude API: %w", err)
	}

	// 解析响应
	keywords, err := parseKeywordsResponse(response, opts.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &KeywordsResult{
		Keywords: keywords,
		Count:    len(keywords),
		Language: opts.Language,
	}, nil
}
