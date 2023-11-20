package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"gopkg.in/yaml.v3"
	"os"
	// "net/http"
	// "time"
	// "webcrawl/config"
	// "./models"
	// "./utils"
)

func main() {
	configData, err := os.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	type Config struct {
		TargetURL string `yaml:"target_url"`
		MaxDepth  int    `yaml:"max_depth"`
	}
	// 將 configData 解析為 YAML 格式的結構體
	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用 config 結構體
	fmt.Println(config.TargetURL) // https://www.example.com
	fmt.Println(config.MaxDepth)  // 2

	// Load configuration from config.yaml file
	// config, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatal("Error loading config:", err)
	// }

	// // Create a new crawler instance
	// crawler := utils.NewCrawler(config)

	// // Start crawling the target URL
	// err = crawler.Crawl(config.TargetURL)
	// if err != nil {
	// 	log.Fatal("Error crawling:", err)
	// }

	// // Print the crawled items
	// for _, item := range crawler.Items {
	// 	fmt.Printf("%s - %s\n", item.Title, item.URL)
	// }
}
