package common

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html"
	"io"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Response struct {
	Header     http.Header
	Body       []byte
	StatusCode string
	Title      string
	Server     string
	Status     string
}

func Request(url string) (*Response, error) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败, %v", err)
	}

	// 跳过证书验证
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: transport,
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("请求失败, %v", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("无法读取响应内容, %v", err)
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("无法读取 Document 内容, %v ", err)
	}

	title := document.Find("title").Text()

	limit := int64(512)

	// limitReader := io.LimitReader(response.Body, limit)

	if int64(len(body)) > limit {
		body = body[:limit]
	}

	status := response.Status
	statusCode := strconv.Itoa(response.StatusCode)
	server := response.Header.Get("Server")

	bodyEscape := html.EscapeString(string(body))

	return &Response{
		Header:     response.Header,
		Body:       []byte(bodyEscape),
		StatusCode: statusCode,
		Title:      title,
		Server:     server,
		Status:     status,
	}, nil

}
