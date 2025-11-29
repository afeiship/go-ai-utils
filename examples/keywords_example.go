package main

import (
	"fmt"
	"log"

	"github.com/afeiship/go-ai-utils"
)

func main() {
	// 示例文本
	content := "人工智能是计算机科学的一个分支，它企图了解智能的实质，并生产出一种新的能以人类智能相似的方式做出反应的智能机器。该领域的研究包括机器人、语言识别、图像识别、自然语言处理和专家系统等。人工智能从诞生以来，理论和技术日益成熟，应用领域也不断扩大。"

	// 方法1: 使用环境变量创建Client (推荐)
	// 设置环境变量: export ANTHROPIC_AUTH_TOKEN="your-api-key"
	fmt.Println("=== 使用环境变量创建Client ===")
	client := aiutils.NewClient(aiutils.NewClientOptions())

	// 默认配置
	result, err := client.Keywords(content)
	if err != nil {
		log.Printf("环境变量方法错误: %v (请确保设置了 ANTHROPIC_AUTH_TOKEN 环境变量)", err)
	} else {
		fmt.Printf("提取到 %d 个关键词:\n", result.Count)
		for i, keyword := range result.Keywords {
			fmt.Printf("%d. %s\n", i+1, keyword)
		}
	}
}
