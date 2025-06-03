# 简易 WAV 文件头部查看器 (Simple WAV Header Inspector)

这是一个使用 Go 语言编写的简单命令行工具，用于查看 `.wav` 音频文件的基本头部信息。

## 功能

* 检查文件是否包含有效的 "RIFF"、"WAVE" 和 "fmt " 标记。
* 显示音频格式和声道数的原始字节值（来自文件头部的特定偏移量）。
* 
## 示例

假设你有一个名为 `example.wav` 的文件：

```bash
go run main.go example.wav
```