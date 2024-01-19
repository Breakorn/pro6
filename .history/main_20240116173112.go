package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mbndr/logo"
)

func main() {

	log := logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)
	// file, err := os.Stat("./db/test.txt") // For read access.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 打开文本文件
	file, err := os.Open("./db/test.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个带有默认缓冲区的读取器
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		// 处理每一行的内容，这里只是简单地打印出来
		fmt.Println(line)
	}

	// 检查扫描过程是否出现错误
	if err := scanner.Err(); err != nil {
		fmt.Println("扫描文件时发生错误:", err)
	}
	fmt.Printf("file: %v\n", file)
	log.Debug("First debug", " and another string to log")
	log.Info("Information")
	log.Warn("Warning", " message")
	log.Error("Error message")
	log.Fatal("Fatal error", " because of something")

	// log.Debugf("Debug value %d", 16)
	// log.Infof("Listening on port %d", 8080)

}
