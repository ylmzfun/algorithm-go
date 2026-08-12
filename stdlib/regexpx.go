package stdlib

// 思路：regexp 基于 RE2 语法，将正则表达式编译为内部自动机后进行匹配。
// 作用：提供文本模式匹配、捕获分组提取与替换。
// 业务场景：
// 1. 表单校验（邮箱、手机号、身份证号）
// 2. 日志解析（从日志行中提取 IP、时间戳）
// 3. 敏感信息脱敏（手机号、银行卡号打码）

import (
	"fmt"
	"regexp"
)

// emailPattern 邮箱正则（预编译，避免每次匹配都重新编译）
var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// emailInTextPattern 用于从文本中提取邮箱（不带 ^$ 锚点）
var emailInTextPattern = regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)

// phonePattern 中国大陆手机号正则：1 开头 + 第二位 3-9 + 9 位数字
var phonePattern = regexp.MustCompile(`^1[3-9]\d{9}$`)

// datePattern 日期正则（命名分组捕获年/月/日）
var datePattern = regexp.MustCompile(`^(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})$`)

// phoneMaskPattern 手机号打码正则：保留前 3 位与后 4 位
var phoneMaskPattern = regexp.MustCompile(`(\d{3})\d{4}(\d{4})`)

// IsValidEmail 校验邮箱格式
func IsValidEmail(s string) bool {
	return emailPattern.MatchString(s)
}

// IsValidPhone 校验手机号格式
func IsValidPhone(s string) bool {
	return phonePattern.MatchString(s)
}

// ExtractEmails 从文本中提取所有邮箱地址
func ExtractEmails(s string) []string {
	return emailInTextPattern.FindAllString(s, -1)
}

// ExtractDateParts 从形如 "2024-01-15" 的日期字符串中提取年/月/日（命名分组）
// 返回是否匹配成功
func ExtractDateParts(s string) (year, month, day string, ok bool) {
	m := datePattern.FindStringSubmatch(s)
	if m == nil {
		return "", "", "", false
	}
	return m[datePattern.SubexpIndex("year")],
		m[datePattern.SubexpIndex("month")],
		m[datePattern.SubexpIndex("day")], true
}

// MaskPhone 手机号脱敏：将中间 4 位替换为 ****
// "13812345678" -> "138****5678"
func MaskPhone(s string) string {
	return phoneMaskPattern.ReplaceAllString(s, "${1}****${2}")
}

// CompilePattern 编译正则表达式，返回错误而非 panic
// 与 regexp.MustCompile 不同，非法表达式不会导致程序崩溃
func CompilePattern(pattern string) (*regexp.Regexp, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("compile regexp %q: %w", pattern, err)
	}
	return re, nil
}
