package core

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"webscr/common"
	"webscr/configs"

	"github.com/chromedp/chromedp"
)

// type RequestErr struct {
// 	errLogs string
// }

func screenshot(url <-chan string, output string, retrys, waitTime, extraTime int) {

	// colorRedPrint := color.New()
	// colorRedPrint.Add(color.FgRed)

	colorRedPrint := common.Colors(common.ColorRed)

	var errLogs []string

	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", true), // 忽略证书错误
		chromedp.Flag("disable-web-security", true),      // 禁用安全策略
	)

	// 自定义截图尺寸
	width := 1920
	height := 1080

	for urls := range url {

		func() {
			execAll, cancel := chromedp.NewExecAllocator(context.Background(), options...)
			defer cancel()

			cont, cancel := chromedp.NewContext(execAll)
			defer cancel()

			wait := waitTime
			ctx, cancel := context.WithTimeout(cont, time.Second*time.Duration(wait))
			defer cancel()

			var buf []byte

			maxRetry := retrys
			extra := extraTime

			var err error

			for retry := 0; retry < maxRetry; retry++ {
				err = chromedp.Run(ctx,
					chromedp.EmulateViewport(int64(width), int64(height)),
					chromedp.Navigate(urls),
					chromedp.WaitReady("body"),
					chromedp.Sleep(time.Second*time.Duration(extra)),
					chromedp.CaptureScreenshot(&buf),
				)

				if err == nil {
					break
				}

				if strings.Contains(err.Error(), "ERR_CONNECTION_TIMED_OUT") || strings.Contains(err.Error(), "ERR_CONNECTION_RESET") {
					time.Sleep(time.Second * 3)
					continue
				}

				errLogs = append(errLogs, urls)

				colorRedPrint.Printf("无法截取 %v, %v\n", urls, err)
				return
			}

			if err != nil {

				errLogs = append(errLogs, urls)

				colorRedPrint.Printf("无法截取 %v, %v\n", urls, err)
				return
			}

			urlRap := strings.ReplaceAll(urls, "http://", "http_")
			urlRap = strings.ReplaceAll(urlRap, "https://", "https_")
			urlRap = strings.ReplaceAll(urlRap, ":", "%")
			urlRap = strings.ReplaceAll(urlRap, "/", "_") // 如果域名后面无路径可设置为 ""
			fileName := fmt.Sprintf("%s/%v.png", output, urlRap)

			if err := os.WriteFile(fileName, buf, 0644); err != nil {
				colorRedPrint.Printf("无法保存截图 %v, %v\n", urls, err)
				return
			}

			response, err := common.Request(urls)
			if err != nil {
				fmt.Println(err)
			}

			// re参数为true, 打印响应信息, 并保存
			if configs.Requests {
				urlsSplit := strings.Split(urls, "/")

				address := common.ParseIP(urlsSplit[2])

				colorGreenPrint := common.Colors(common.ColorGreen)

				statusCode := response.StatusCode
				title := response.Title
				server := response.Server

				colorGreenPrint.Printf("%v %v %v %v %v\n", urls, address, title, statusCode, server)

				common.WriteCsv(urls, address, title, statusCode, server)
			} else {
				colorGreenPrint := common.Colors(common.ColorGreen)
				colorGreenPrint.Printf("%v \n", urls)
			}

		}()
	}

	// 记录因上下文超时或其他原因导致截图失败的URL, 方便后续操作
	common.Write(errLogs)
}

func Runscreenshot() {
	configs.Flag()
	flag.Parse()

	// colorYellowPrint := color.New()
	// colorYellowPrint.Add(color.FgYellow)

	colorYellowPrint := common.Colors(common.ColorYellow)

	if configs.UrlsFiles == "" {
		fmt.Println("-f 参数不能为空, --help 查看帮助")
		os.Exit(1)
	}

	openfile, err := os.Open(configs.UrlsFiles)
	if err != nil {
		fmt.Println("无法打开文件, ", err)
		return
	}

	defer openfile.Close()

	scanner := bufio.NewScanner(openfile)

	if err := os.MkdirAll(configs.OutputDir, 0755); err != nil {
		fmt.Printf("无法创建 %v 目录, %v\n", configs.OutputDir, err)
		return
	}

	var urlDe = make(map[string]struct{})

	for scanner.Scan() {
		if scanner.Text() != "" {
			// 去重
			urlDe[scanner.Text()] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取失败, ", err)
		return
	}

	var urlChan = make(chan string, max(100, len(urlDe)))

	var wg sync.WaitGroup

	for i := 0; i < configs.Thread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			screenshot(urlChan, configs.OutputDir, configs.Retrys, configs.WaitTime, configs.ExtraTime)
		}()
	}

	go func() {
		for urlValue := range urlDe {
			urlChan <- urlValue
		}
		close(urlChan)
	}()

	wg.Wait()

	if configs.OutputHtml != "" {
		colorYellowPrint.Printf("正在输出 %v 文件\n", configs.OutputHtml)
		common.WriteHtml()
	}

	colorYellowPrint.Println("所有任务已完成!")
}
