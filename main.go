package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 检查是否给了文件名
	if len(os.Args) < 2 {
		fmt.Println("请提供一个 WAV 文件名作为参数")
		fmt.Println("例如: go run main.go audio.wav")
		return
	}
	filePath := os.Args[1]

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("打开文件失败:", err) // 如果打开失败，就直接退出
	}
	defer file.Close() // main函数结束前关闭文件

	// 读取文件头的前36个字节
	header := make([]byte, 36)
	_, err = io.ReadFull(file, header) // 尝试读取36个字节
	if err != nil {
		log.Fatal("读取文件头失败:", err) // 如果读取不够36字节或出错，也退出
	}

	fmt.Println("--- WAV 文件头部信息 ---")
	fmt.Printf("文件名: %s\n\n", filePath)

	// 检查 "RIFF"
	riffChunkID := string(header[0:4])
	if riffChunkID == "RIFF" {
		fmt.Println("RIFF 标记: 是")
	} else {
		fmt.Printf("RIFF 标记: 否 (找到: %s)\n", riffChunkID)
	}

	// 检查 "WAVE"
	waveFormat := string(header[8:12])
	if waveFormat == "WAVE" {
		fmt.Println("WAVE 格式: 是")
	} else {
		fmt.Printf("WAVE 格式: 否 (找到: %s)\n", waveFormat)
	}

	// 检查 "fmt "
	fmtSubChunkID := string(header[12:16])
	if fmtSubChunkID == "fmt " {
		fmt.Println("fmt  标记: 是")
	} else {
		fmt.Printf("fmt  标记: 否 (找到: %s)\n", fmtSubChunkID)
	}

	// 显示音频格式和声道数的原始字节 (新手可能这么做)
	// AudioFormat 在偏移量 20 (2字节)
	// NumChannels 在偏移量 22 (2字节)
	audioFormatLowByte := header[20]
	audioFormatHighByte := header[21]
	fmt.Printf("音频格式 (字节@20,21): %d, %d  (通常 1,0 代表 PCM)\n", audioFormatLowByte, audioFormatHighByte)

	numChannelsLowByte := header[22]
	numChannelsHighByte := header[23]
	fmt.Printf("声道数 (字节@22,23): %d, %d  (通常 1,0 代表单声道, 2,0 代表立体声)\n", numChannelsLowByte, numChannelsHighByte)

	fmt.Println("\n提示: 以上只是原始字节的十进制值。")
}
