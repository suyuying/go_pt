package main

import (
	"fmt"
	// "github.com/gocolly/colly/v2"
	"webcrawl/config"
)

func main() {
	// Load configuration from config.yaml file
	configss:= config.LoadConfig("config/config.yaml")
	fmt.Println(configss)
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
