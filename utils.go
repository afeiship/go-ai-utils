package aiutils

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/afeiship/go-claude"
	"gopkg.in/yaml.v3"
)

//go:embed assets/strings.yml
var stringsFile embed.FS

// StringsConfig 字符串配置文件结构
type StringsConfig struct {
	Prompts struct {
		ChineseKeywords string `yaml:"chinese_keywords"`
		EnglishKeywords string `yaml:"english_keywords"`
		MixedKeywords   string `yaml:"mixed_keywords"`
	} `yaml:"prompts"`
	LanguageInstructions map[string]string `yaml:"language_instructions"`
	Defaults struct {
		Count     int    `yaml:"count"`
		Language  string `yaml:"language"`
		Model     string `yaml:"model"`
		MaxTokens int    `yaml:"max_tokens"`
		BaseURL   string `yaml:"base_url"`
	} `yaml:"defaults"`
}

// PromptData 提示词模板数据
type PromptData struct {
	Count   int
	Content string
}

var (
	// 全局字符串配置
	stringConfig *StringsConfig
	// 提示词模板缓存
	promptTemplates = make(map[string]*template.Template)
)

// init 初始化字符串配置
func init() {
	// 读取嵌入的配置文件
	data, err := stringsFile.ReadFile("assets/strings.yml")
	if err != nil {
		panic(fmt.Sprintf("failed to read assets/strings.yml: %v", err))
	}

	// 解析YAML配置
	var config StringsConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(fmt.Sprintf("failed to parse assets/strings.yml: %v", err))
	}

	stringConfig = &config

	// 预编译提示词模板
	if err := compilePromptTemplates(); err != nil {
		panic(fmt.Sprintf("failed to compile prompt templates: %v", err))
	}
}

// getStringConfig 获取字符串配置
func getStringConfig() *StringsConfig {
	if stringConfig == nil {
		panic("stringConfig not initialized")
	}
	return stringConfig
}

// compilePromptTemplates 编译提示词模板
func compilePromptTemplates() error {
	templates := map[string]string{
		"chinese": stringConfig.Prompts.ChineseKeywords,
		"english": stringConfig.Prompts.EnglishKeywords,
		"mixed":   stringConfig.Prompts.MixedKeywords,
	}

	for name, content := range templates {
		tmpl, err := template.New(name).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse %s template: %w", name, err)
		}
		promptTemplates[name] = tmpl
	}

	return nil
}

// buildPrompt 构建AI提示词
func buildPrompt(content string, language Language, count int) (string, error) {
	// 根据语言选择对应的模板
	templateName := string(language)
	tmpl, exists := promptTemplates[templateName]
	if !exists {
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	// 准备模板数据
	data := PromptData{
		Count:   count,
		Content: content,
	}

	// 执行模板
	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// parseKeywordsResponse 解析Claude响应
func parseKeywordsResponse(response string, count int) ([]string, error) {
	// 清理响应文本
	response = strings.TrimSpace(response)
	lines := strings.Split(response, "\n")

	var keywords []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			// 移除可能的序号和符号
			line = strings.TrimPrefix(line, "-")
			line = strings.TrimPrefix(line, "*")
			line = strings.TrimPrefix(line, "•")
			line = strings.TrimPrefix(line, fmt.Sprintf("%d.", len(keywords)+1))
			line = strings.TrimSpace(line)

			if line != "" {
				keywords = append(keywords, line)
			}
		}
	}

	if len(keywords) == 0 {
		return nil, fmt.Errorf("no keywords found in response")
	}

	// 限制关键词数量
	if len(keywords) > count {
		keywords = keywords[:count]
	}

	return keywords, nil
}

// createClaudeClient 创建Claude客户端
func createClaudeClient(apiKey, baseURL, model string, maxTokens int) (*claude.Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	return claude.NewClient(claude.Config{
		APIKey:    apiKey,
		BaseURL:   baseURL,
		Model:     model,
		MaxTokens: maxTokens,
	})
}