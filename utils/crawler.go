package utils

//https://go-colly.org/docs/
import (
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	// "github.com/gocolly/colly/debug"
)

func Crawl(domain string) {
	// 物件入口
	c := colly.NewCollector()
	// Add the headers
	// Set the headers
	c.SetCookies(domain, []*http.Cookie{
		{
			Name:  "over18",
			Value: "1",
		},
	})

	// 設定 Debug 選項
	// 設置 Async 選項為 true
	// c.Async = true
	// 設定爬蟲相關call back function

	// 執行操作 先選指定html tag,選完觸發call back function
	// 第二個參數是匿名函示
	// c.OnHTML("a[href]", func(h *colly.HTMLElement) {
	// 	link := h.Attr("href")
	// 	fmt.Printf("Link found: %q -> %s\n", h.Text, link)
	// 	// c.Visit(h.Request.AbsoluteURL(link))
	// })
	// 註冊一個 OnHTML 事件處理器
	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// 註冊一個 OnError 事件處理器
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("發生錯誤：", err)
	})
	c.Visit(domain)
	// c.Wait()
}
