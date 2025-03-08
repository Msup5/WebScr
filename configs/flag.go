package configs

import "flag"

func Flag() {
	flag.StringVar(&UrlsFiles, "f", "", "URL 文件路径")
	flag.StringVar(&OutputDir, "o", "image", "保存路径")
	flag.IntVar(&Retrys, "r", 3, "重试次数")
	flag.IntVar(&WaitTime, "wT", 60, "等待时间")
	flag.IntVar(&ExtraTime, "eT", 19, "额外等待时间")
	flag.IntVar(&Thread, "t", 5, "线程")
	flag.BoolVar(&Requests, "re", false, "获取标题, 状态码")
}
