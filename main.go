package main

import (
	"fmt"
	// "github.com/gocolly/colly/v2"
	"webcrawl/config"
	"webcrawl/utils"
)

type Website struct {
	TargetURL   string `yaml:"target_url"`
	Description string `yaml:"description"`
}
type Config struct {
	Websites []Website `yaml:"websites"`
}

func main() {
	var config_web config.Config
	// Load configuration from config.yaml file
	err := config.LoadConfig("config/config.yaml", &config_web)
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}
	for element_number, website := range config_web.Websites {
		fmt.Printf("URL: %s, Description: %s,%d\n", website.TargetURL, website.Description, element_number)
		utils.Crawl(website.TargetURL)
	}

}
