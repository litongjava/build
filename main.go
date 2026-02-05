package main

import (
	"log"
	"os"
)

func init() {
	// 设置Flags为 日期 时间 微秒 文件名:行号
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	var filePath = ""
	if len(os.Args) >= 2 {
		filePath = os.Args[1]
	} else {
		filePath = findBuildFile()
	}

	Build(filePath)
}

func findBuildFile() string {
	files := []string{".build.txt", "build.txt"}
	for _, f := range files {
		if _, err := os.Stat(f); err == nil {
			return f
		}
	}
	log.Fatal("no build file found (.build.txt or build.txt)")
	return ""
}
