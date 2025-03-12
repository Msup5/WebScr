package common

import (
	"fmt"
	"os"
)

func Write(writeFile []string) {
	file, err := os.OpenFile("tmp/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		fmt.Println("无法打开文件, ", err)
		return
	}

	defer file.Close()

	for _, cont_value := range writeFile {
		if _, err := file.WriteString(cont_value + "\n"); err != nil {
			fmt.Println("写入文件失败, ", err)
			return
		}
	}

}
