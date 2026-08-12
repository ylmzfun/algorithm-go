package stdlib

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLineReader(t *testing.T) {
	lr := NewLineReader(strings.NewReader("line1\nline2\nline3\n"))
	var lines []string
	for {
		line, ok := lr.Next()
		if !ok {
			break
		}
		lines = append(lines, line)
	}
	if lr.Err() != nil {
		t.Errorf("Expected no scan error, got %v", lr.Err())
	}
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "line1" || lines[2] != "line3" {
		t.Errorf("Expected line1/line3, got %s/%s", lines[0], lines[2])
	}
}

func TestLineReaderEmptyInput(t *testing.T) {
	lr := NewLineReader(strings.NewReader(""))
	line, ok := lr.Next()
	if ok || line != "" {
		t.Errorf("Expected EOF with empty line, got %q/%v", line, ok)
	}
}

func TestLineReaderTooLongLine(t *testing.T) {
	// 单行超过 bufio.Scanner 默认 64KB 上限时返回错误
	long := strings.Repeat("a", 70*1024)
	lr := NewLineReader(strings.NewReader(long))
	if _, ok := lr.Next(); ok {
		t.Error("Expected scan failure for over-long line")
	}
	if lr.Err() == nil {
		t.Error("Expected error for over-long line")
	}
}

func TestReadUntil(t *testing.T) {
	s, err := ReadUntil(strings.NewReader("a;b;c"), ';')
	if err != nil {
		t.Fatalf("ReadUntil failed: %v", err)
	}
	if s != "a;" {
		t.Errorf("Expected \"a;\", got %q", s)
	}

	// 源数据结束仍返回已读内容
	s, err = ReadUntil(strings.NewReader("no-delim"), ';')
	if err != nil {
		t.Errorf("Expected nil error at EOF with data, got %v", err)
	}
	if s != "no-delim" {
		t.Errorf("Expected \"no-delim\", got %q", s)
	}

	// 空输入返回 io.EOF
	if _, err = ReadUntil(strings.NewReader(""), ';'); err != io.EOF {
		t.Errorf("Expected io.EOF, got %v", err)
	}
}

func TestWriteLines(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteLines(&buf, []string{"a", "b", "c"}); err != nil {
		t.Fatalf("WriteLines failed: %v", err)
	}
	if buf.String() != "a\nb\nc\n" {
		t.Errorf("Expected \"a\\nb\\nc\\n\", got %q", buf.String())
	}
}

func TestCopyToWriter(t *testing.T) {
	var buf bytes.Buffer
	src := "hello, 标准库 io"
	n, err := CopyToWriter(&buf, strings.NewReader(src))
	if err != nil {
		t.Fatalf("CopyToWriter failed: %v", err)
	}
	if n != int64(len(src)) {
		t.Errorf("Expected %d bytes copied, got %d", len(src), n)
	}
	if buf.String() != src {
		t.Errorf("Expected %q, got %q", src, buf.String())
	}
}

func TestCountWords(t *testing.T) {
	count, err := CountWords("hello  world\nfoo bar")
	if err != nil {
		t.Fatalf("CountWords failed: %v", err)
	}
	if count != 4 {
		t.Errorf("Expected 4 words, got %d", count)
	}
}
