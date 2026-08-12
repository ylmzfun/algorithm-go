package stdlib

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	valid := []string{
		"user@example.com",
		"a.b+c@sub.example.co.uk",
		"name_123@test.io",
	}
	for _, email := range valid {
		if !IsValidEmail(email) {
			t.Errorf("Expected %q to be valid", email)
		}
	}

	invalid := []string{
		"plainaddress",
		"a@b",             // 顶级域名过短
		"@example.com",    // 缺少用户名
		"user name@x.com", // 含空格
	}
	for _, email := range invalid {
		if IsValidEmail(email) {
			t.Errorf("Expected %q to be invalid", email)
		}
	}
}

func TestIsValidPhone(t *testing.T) {
	if !IsValidPhone("13812345678") {
		t.Error("Expected 13812345678 to be valid")
	}
	for _, phone := range []string{"12345678901", "23812345678", "1381234567"} {
		if IsValidPhone(phone) {
			t.Errorf("Expected %q to be invalid", phone)
		}
	}
}

func TestExtractEmails(t *testing.T) {
	text := "请联系 support@example.com 或 admin@test.org，非邮箱 xyz@bad"
	emails := ExtractEmails(text)
	if len(emails) != 2 {
		t.Fatalf("Expected 2 emails, got %v", emails)
	}
	if emails[0] != "support@example.com" || emails[1] != "admin@test.org" {
		t.Errorf("Unexpected emails: %v", emails)
	}
}

func TestExtractDateParts(t *testing.T) {
	year, month, day, ok := ExtractDateParts("2024-01-15")
	if !ok {
		t.Fatal("Expected date to match")
	}
	if year != "2024" || month != "01" || day != "15" {
		t.Errorf("Expected 2024/01/15, got %s/%s/%s", year, month, day)
	}

	if _, _, _, ok := ExtractDateParts("2024/01/15"); ok {
		t.Error("Expected no match for slash-separated date")
	}
	if _, _, _, ok := ExtractDateParts("not-a-date"); ok {
		t.Error("Expected no match for invalid date")
	}
}

func TestMaskPhone(t *testing.T) {
	if got := MaskPhone("13812345678"); got != "138****5678" {
		t.Errorf("Expected \"138****5678\", got %q", got)
	}
}

func TestCompilePattern(t *testing.T) {
	// 合法表达式可正常编译并使用
	re, err := CompilePattern(`^\d{4}$`)
	if err != nil {
		t.Fatalf("CompilePattern failed: %v", err)
	}
	if !re.MatchString("2024") {
		t.Error("Expected 2024 to match ^\\d{4}$")
	}
	if re.MatchString("24") {
		t.Error("Expected 24 not to match ^\\d{4}$")
	}

	// 非法表达式返回错误而不是 panic
	if _, err := CompilePattern(`(unclosed`); err == nil {
		t.Error("Expected compile error for invalid pattern")
	}
}
