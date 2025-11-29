package main

import (
	"fmt"
	"log"

	"github.com/afeiship/go-ai-utils"
)

func main() {
	content := "人工智能是计算机科学的一个分支"

	// 测试重构后的API
	client := aiutils.NewClient(aiutils.ClientOptions{})

	result, err := client.Keywords(content, &aiutils.KeywordsOptions{
		Count:    3,
		Language: aiutils.LanguageChinese,
	})

	if err != nil {
		log.Fatalf("错误: %v", err)
	}

	fmt.Printf("✅ 重构成功！提取到 %d 个关键词:\n", result.Count)
	for i, keyword := range result.Keywords {
		fmt.Printf("%d. %s\n", i+1, keyword)
	}
}
