package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	urlsFile, outputDir string
	threads             int
)

func init() {
	flag.StringVar(&urlsFile, "u", "", "URL 文件路径")
	flag.StringVar(&outputDir, "o", "image", "保存路径")
	flag.IntVar(&threads, "t", 5, "线程")
}

func screenshot(url <-chan string, output string) {
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", true), // 忽略证书错误
		chromedp.Flag("disable-web-security", true),      // 禁用安全策略
	)

	for urls := range url {

		func() {
			execAll, cancel := chromedp.NewExecAllocator(context.Background(), options...)
			defer cancel()

			cont, cancel := chromedp.NewContext(execAll)
			defer cancel()

			ctx, cancel := context.WithTimeout(cont, time.Second*60)
			defer cancel()

			var buf []byte

			maxRetry := 3 // 遇到无法解析域名或被重置连接, 重试最多3次
			var err error

			for retry := 0; retry < maxRetry; retry++ {
				err = chromedp.Run(ctx,
					chromedp.Navigate(urls),
					chromedp.WaitReady("body"),
					chromedp.Sleep(time.Second*10),
					chromedp.FullScreenshot(&buf, 90),
				)

				if err == nil {
					break
				}

				if strings.Contains(err.Error(), "ERR_CONNECTION_TIMED_OUT") || strings.Contains(err.Error(), "ERR_CONNECTION_RESET") {
					time.Sleep(time.Second * 3)
					continue
				}

				fmt.Printf("无法截取 %v, %v\n", urls, err)
				return
			}
			if err != nil {
				fmt.Printf("无法截取 %v, %v\n", urls, err)
				return
			}

			urlRap := strings.ReplaceAll(urls, "http://", "http_")
			urlRap = strings.ReplaceAll(urlRap, "https://", "https_")
			urlRap = strings.ReplaceAll(urlRap, ":", "_")
			urlRap = strings.ReplaceAll(urlRap, "/", "")
			fileName := fmt.Sprintf("%s/%v.png", output, urlRap)

			if err := os.WriteFile(fileName, buf, 0644); err != nil {
				fmt.Printf("无法保存截图 %v, %v\n", urls, err)
				return
			}

		}()
	}
}

func main() {
	flag.Parse()

	if urlsFile == "" {
		fmt.Println(`
 _       __     __   _____          
| |     / /__  / /_ / ___/__________
| | /| / / _ \/ __ \\__ \/ ___/ ___/
| |/ |/ /  __/ /_/ /__/ / /__/ /    
|__/|__/\___/_.___/____/\___/_/ v1.0.3
		
参数:
  -u  URL 文件路径 (必须)
  -o  保存路径
  -t  线程
		`)
		return
	}

	// outputDir := "image"

	openFile, err := os.Open(urlsFile)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	scanner := bufio.NewScanner(openFile)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("无法创建 %v 目录, %v\n", outputDir, err)
	}

	var urlDe = make(map[string]struct{})

	for scanner.Scan() {
		if scanner.Text() != "" {
			// 去重
			urlDe[scanner.Text()] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取失败", err)
		return
	}

	var urlChan = make(chan string, max(100, len(urlDe)))

	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			screenshot(urlChan, outputDir)
		}()
	}

	go func() {
		for urlValue := range urlDe {
			urlChan <- urlValue
		}
		close(urlChan)
	}()

	wg.Wait()
	fmt.Println("已完成所有截图")

}
