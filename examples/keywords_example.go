package main

import (
	"context"
	"fmt"
	"log"

	"github.com/afeiship/aiutils"
)

func main() {
	ctx := context.Background()

	// 示例文本
	content := "人工智能是计算机科学的一个分支，它企图了解智能的实质，并生产出一种新的能以人类智能相似的方式做出反应的智能机器。该领域的研究包括机器人、语言识别、图像识别、自然语言处理和专家系统等。人工智能从诞生以来，理论和技术日益成熟，应用领域也不断扩大。"

	// 方法1: 使用环境变量创建Client (推荐)
	// 设置环境变量: export ANTHROPIC_AUTH_TOKEN="your-api-key"
	fmt.Println("=== 使用环境变量创建Client ===")
	client := aiutils.NewClientFromEnv()
	keywords, err := client.KeywordsSimple(ctx, content)
	if err != nil {
		log.Printf("环境变量方法错误: %v (请确保设置了 ANTHROPIC_AUTH_TOKEN 环境变量)", err)
	} else {
		fmt.Println("提取的关键词:")
		for i, keyword := range keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}

	// 方法2: 直接传入API Key创建Client
	fmt.Println("\n=== 直接传入API Key ===")
	apiKey := "your-api-key" // 请替换为你的Claude API密钥
	if apiKey != "your-api-key" { // 只有在用户修改了API key时才执行
		clientWithKey := aiutils.NewClient(apiKey)
		keywords, err = clientWithKey.KeywordsSimple(ctx, content)
		if err != nil {
			log.Printf("直接传入API Key错误: %v", err)
		} else {
			fmt.Println("提取的关键词:")
			for i, keyword := range keywords {
				fmt.Printf("%d. %s\n", i+1, keyword)
			}
		}
	} else {
		fmt.Println("请修改示例代码中的API密钥后再测试")
	}

	// 方法3: 指定关键词数量
	fmt.Println("\n=== 指定关键词数量 ===")
	keywords3, err := client.KeywordsWithCount(ctx, content, 3)
	if err != nil {
		log.Printf("指定数量错误: %v", err)
	} else {
		fmt.Println("提取3个关键词:")
		for i, keyword := range keywords3 {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}

	// 方法4: 指定语言
	fmt.Println("\n=== 指定语言 - 英文 ===")
	keywordsEn, err := client.KeywordsWithLanguage(ctx, content, aiutils.LanguageEnglish)
	if err != nil {
		log.Printf("英文关键词错误: %v", err)
	} else {
		fmt.Println("提取英文关键词:")
		for i, keyword := range keywordsEn {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}

	// 方法5: 使用完整选项
	fmt.Println("\n=== 完整选项配置 ===")
	result, err := client.Keywords(ctx, content, &aiutils.KeywordsOptions{
		Count:    8,
		Language: aiutils.LanguageMixed,
	})
	if err != nil {
		log.Printf("完整选项错误: %v", err)
	} else {
		fmt.Printf("提取到 %d 个混合语言关键词:\n", result.Count)
		for i, keyword := range result.Keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}

	// 方法6: 链式配置Client
	fmt.Println("\n=== 链式配置Client ===")
	chainedClient := aiutils.NewClientFromEnv().
		SetModel("glm-4.5-air").
		SetMaxTokens(1500)

	keywords, err = chainedClient.KeywordsWithCount(ctx, content, 5)
	if err != nil {
		log.Printf("链式配置错误: %v", err)
	} else {
		fmt.Println("链式配置提取的关键词:")
		for i, keyword := range keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}

	// 方法7: 创建自定义Client
	fmt.Println("\n=== 自定义Client配置 ===")
	customClient := aiutils.NewClientWithOptions(
		"", // 空字符串表示使用环境变量
		"https://api.anthropic.com",
		"glm-4.5-air",
		2000,
	)

	keywords, err = customClient.KeywordsWithLanguage(ctx, content, aiutils.LanguageChinese)
	if err != nil {
		log.Printf("自定义Client错误: %v", err)
	} else {
		fmt.Println("自定义Client提取的中文关键词:")
		for i, keyword := range keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}
}