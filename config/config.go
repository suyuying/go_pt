package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"fmt"
)
//var configInstance Config
// configPointer := &configInstance
// configValue := *configPointer
// 在声明时，* 表示指针类型：
// var p *int
// 在这里，p 是一个指向整数的指针。这是在声明指针变量时使用 * 的一种情况。

// 在使用时，* 用于解引用：
// var x int = 42
// var p *int = &x
// fmt.Println(*p)
// 在这里，*p 表示获取指针 p 所指向的值。这是在使用指针时使用 * 的一种情况。
// 這格式是配合yaml.v3去做結構標籤,告訴yaml如何解讀對應key
type Website struct {
	TargetURL string `yaml:"target_url"`
	Description string `yaml:"description"`
}
type Config struct {
	Websites []Website `yaml:"websites"`
}

func LoadConfig(route string) (config Config){
	configData, err := os.ReadFile(route)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 將 configData 解析為 YAML 格式的結構體
	// 把數據讀到解析的結構體中
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	return config
	// 使用 config 結構體

}