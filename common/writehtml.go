package common

import (
	"fmt"
	"os"
	"strings"
	"webscr/configs"

	"github.com/klarkxy/gohtml"
)

// func work(filename os.DirEntry) {

// }

func WriteHtml() {
	newhtm := gohtml.NewHtml()

	newhtm.Meta().Charset("utf-8")

	newhtm.Head().Title().Text("WebScr 执行结果")

	newhtm.Head().Style().Text(`
	table {
		border-collapse: collapse;	
		table-layout:fixed;
    }

    table, th, td {
		border: 1px solid rgb(203, 195, 195);
		width:100%;
		height: 30px;
		word-break:keep-all;
		white-space:nowrap;
		overflow:hidden;
		text-overflow:ellipsis;
    }

	#responseTd {
		vertical-align: top;
		text-align: left;
		font-size: 12px;
		white-space: pre-wrap;
	}
	`)

	// 输出 html 文件的路径
	outputTtmlDir := configs.OutputHtml

	// imageDir := configs.OutputDir

	htmbody := newhtm.Body().Tag("div").Align("center").Id("content").Tag("table").Align("center").Tag("tbody")

	htmbody.Tr().Align("center").Td().Colspan("5").Text(outputTtmlDir)

	showTr := htmbody.Tr().Align("center").Attr("style", "color:black")
	showTr.Td().Text("URL")
	showTr.Td().Text("IP")
	showTr.Td().Text("状态码")
	showTr.Td().Text("标题")
	showTr.Td().Text("服务器")

	dir := configs.OutputDir

	// fmt.Println(dir)

	dirList, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("打开 %v 文件失败, %v\n", dirList, err)
		return
	}

	for _, filename := range dirList {
		url := InitializeName(filename.Name())

		urlsSplit := strings.Split(url, "/")

		address := ParseIP(urlsSplit[2])

		response, err := Request(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		statusCode := response.StatusCode
		title := response.Title
		server := response.Server

		// body := response.Body

		contTr := htmbody.Tr().Align("center").Attr("style", "color:black")
		contTr.Td().A().Href(url).Text(url)
		contTr.Td().Text(address)
		contTr.Td().Text(statusCode)
		contTr.Td().Text(title)
		contTr.Td().Text(server)

		// img & response 在单独的 tr
		contTrImgResTr := htmbody.Tr().Align("center").Attr("style", "color:black;")
		contTrImgResTr.Td().Colspan("4").Img().Src(dir+"\\"+filename.Name()).Attr("style", "width:800px;hight:450px")

		var headerBuilder strings.Builder

		for key, values := range response.Header {
			for _, value := range values {
				headerBuilder.WriteString(key + ": " + value + "<br>")
			}
		}

		//contTrImgResTr.Td().Text("<br>" + string(response.Body))
		headerBuilder.WriteString("<br>" + string(response.Body) + "<br>")

		contTrImgResTr.Td().Id("responseTd").Text(response.Status + "<br>" + headerBuilder.String())

	}

	htmlOutputDir := configs.OutputHtml

	file, err := os.OpenFile(htmlOutputDir, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("创建 %v 文件失败, %v\n", htmlOutputDir, err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(newhtm.String() + "\n")
	if err != nil {
		fmt.Println("写入数据失败", err)
	}
}
