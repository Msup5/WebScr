package common

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// type RequsetData struct{

// }

func Request(url string) (string, string, string) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("请求失败, ", err)
		return "", "", ""
	}

	client := &http.Client{}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		return "", "", ""
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("无法读取响应内容, ", err)
		return "", "", ""
	}

	statusCode := strconv.Itoa(response.StatusCode)
	title := document.Find("title").Text()
	server := response.Header.Get("Server")

	return statusCode, title, server

}
