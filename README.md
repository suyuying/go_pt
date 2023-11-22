## basic

在Go語言中，使用struct是一種比使用map更常見的方式，差別在靜態
跟動態,大多數都用struct除非是在不確定value的type再用map.

```golang
type Website struct {
 TargetURL string `yaml:"target_url"`
 Description string `yaml:"description"`
}
type Config struct {
 Websites []Website `yaml:"websites"`
}

```

這樣是讀取指定屬性
```golang
# Read
for element_number, website := range config.Websites {
    fmt.Printf("URL: %s, Description: %s,%d\n", website.TargetURL, website.Description, element_number)
}

```

以下方法式操作可以直接對到值在記憶體的情況
```golang
// 這邊要求以指針型態傳入參數,目的是操作對應到變數值記憶體位置
func LoadConfig(route string, configPtr *Config) error {
	// Read the configuration file
    // golang很常這樣安排報錯,把資料跟err放在一起
	configData, err := os.ReadFile(route)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %w", err)
	}

	// Unmarshal the configuration data into the provided Config pointer
    // Unmarshal 部分第二個參數要求船入指標,他會把結果直接改
    // 到對應記體位置
	err = yaml.Unmarshal(configData, configPtr)
	if err != nil {
    // 包裝報錯結果
		return fmt.Errorf("failed to unmarshal configuration data: %w", err)
	}
    // 如果都沒報錯代表有loading到東西
	return nil
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
	}
}
```

另外已很常用到的request套件,他也會操作到指標部分!
&或者*就看他參數帶啥,大部分傳入參數都適用&,在funtion內要使用
真實物件就要用*.
