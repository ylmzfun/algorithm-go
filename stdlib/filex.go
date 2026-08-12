package stdlib

// 思路：os 提供文件/目录的系统级操作，path/filepath 负责跨平台路径拼接、
// 清理与目录遍历。
// 作用：演示文件读写、目录遍历与安全路径拼接。
// 业务场景：
// 1. 配置文件、数据文件的读写
// 2. 日志目录扫描与归档
// 3. 用户上传文件的落盘与路径校验（防路径穿越）

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteFileContent 将文本写入文件（自动创建父目录）
// 文件权限 0644，目录权限 0755
func WriteFileContent(path, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0o644)
}

// ReadFileContent 读取文本文件的完整内容
func ReadFileContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FileExists 判断路径是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// JoinPath 拼接多个路径元素并清理冗余分隔符（filepath.Join 封装）
// 如 JoinPath("a/", "/b/", "../c") -> "a/c"
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

// ListFiles 递归列出 root 目录下的所有文件（不含目录）
// 遍历顺序为字典序
func ListFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// CountFilesByExt 统计 root 目录下指定扩展名（如 ".txt"）的文件数量
// ext 传空串时统计全部文件
func CountFilesByExt(root, ext string) (int, error) {
	files, err := ListFiles(root)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, f := range files {
		if ext == "" || strings.HasSuffix(f, ext) {
			count++
		}
	}
	return count, nil
}

// SafeJoin 拼接 root 与 sub，并确保结果不逃逸出 root
// 用于防止 "../" 路径穿越攻击；sub 为空时返回 root 本身
func SafeJoin(root, sub string) (string, error) {
	clean := filepath.Join(root, sub)
	rel, err := filepath.Rel(root, clean)
	if err != nil {
		return "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("path %q escapes root %q", sub, root)
	}
	return clean, nil
}
