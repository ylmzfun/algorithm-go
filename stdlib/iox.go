package stdlib

// 思路：io 提供读写的最小抽象（Reader/Writer），bufio 通过内存缓冲批量搬运数据、
// 减少系统调用次数，io.Copy 以固定大小缓冲区流式复制任意 Reader 到 Writer。
// 作用：演示缓冲读写、按行扫描与流式复制。
// 业务场景：
// 1. 日志文件按行读取与分析
// 2. 大文件流式复制，避免一次性载入内存
// 3. 网络数据流的透传（代理服务器）

import (
	"bufio"
	"io"
	"strings"
)

// LineReader 按行读取器（基于 bufio.Scanner）
type LineReader struct {
	scanner *bufio.Scanner
}

// NewLineReader 从任意 io.Reader 创建按行读取器
func NewLineReader(r io.Reader) *LineReader {
	return &LineReader{scanner: bufio.NewScanner(r)}
}

// Next 读取下一行内容（不含换行符）
// 返回 false 表示读取完毕；读取失败时同样返回 false，可通过 Err 判断原因
func (lr *LineReader) Next() (string, bool) {
	if !lr.scanner.Scan() {
		return "", false
	}
	return lr.scanner.Text(), true
}

// Err 返回扫描过程中的错误，正常结束时为 nil
// 典型错误：单行超过 bufio.Scanner 默认 64KB 上限
func (lr *LineReader) Err() error {
	return lr.scanner.Err()
}

// ReadUntil 使用 bufio.Reader 读取直到遇到分隔符 delim（结果含分隔符）
// 源数据在遇到分隔符前结束（io.EOF）时返回已读数据与 nil
func ReadUntil(r io.Reader, delim byte) (string, error) {
	line, err := bufio.NewReader(r).ReadString(delim)
	if err == io.EOF && line != "" {
		return line, nil
	}
	return line, err
}

// WriteLines 使用 bufio.Writer 批量写入多行文本
// 写入先进入内存缓冲，最后一次性 Flush 写出，减少底层 Write 调用
func WriteLines(w io.Writer, lines []string) error {
	bw := bufio.NewWriter(w)
	for _, line := range lines {
		if _, err := bw.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return bw.Flush()
}

// CopyToWriter 将 src 内容流式复制到 dst（io.Copy 封装）
// 返回复制的字节数；内部使用 32KB 缓冲，适合大文件/流复制
func CopyToWriter(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

// CountWords 统计文本中的单词个数
// 演示 bufio.Scanner 自定义 SplitFunc（bufio.ScanWords）
func CountWords(text string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
