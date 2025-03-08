package common

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteCsv(url, address, title, statusCode, server string) {
	csvOutput := "results/results.csv"

	fileExists := false

	if fileInfo, err := os.Stat(csvOutput); err == nil {
		fileExists = fileInfo.Size() > 0
	}

	// createCsvFile, err := os.Create(csvOutput)
	createCsvFile, err := os.OpenFile(csvOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		fmt.Println("创建csv文件失败, ", err)
		return
	}

	defer createCsvFile.Close()

	// 解决乱码问题
	if !fileExists {
		if _, err := createCsvFile.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
			return
		}
	}

	writecsv := csv.NewWriter(createCsvFile)
	defer writecsv.Flush()

	if !fileExists {
		header := []string{"URL", "IP", "标题", "状态码", "服务器"}
		if err := writecsv.Write(header); err != nil {
			fmt.Println("写入表头失败:", err)
			return
		}
	}

	data := []string{url, address, title, statusCode, server}

	if err := writecsv.Write(data); err != nil {
		fmt.Println("写入数据失败:", err)
		return
	}

}
