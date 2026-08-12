package stdlib

import (
	"path/filepath"
	"testing"
)

func TestWriteReadFileContent(t *testing.T) {
	dir := t.TempDir()
	// 父目录不存在，WriteFileContent 应自动创建
	path := filepath.Join(dir, "nested", "deep", "a.txt")
	if err := WriteFileContent(path, "hello"); err != nil {
		t.Fatalf("WriteFileContent failed: %v", err)
	}
	content, err := ReadFileContent(path)
	if err != nil {
		t.Fatalf("ReadFileContent failed: %v", err)
	}
	if content != "hello" {
		t.Errorf("Expected \"hello\", got %q", content)
	}
}

func TestReadFileContentNotFound(t *testing.T) {
	dir := t.TempDir()
	if _, err := ReadFileContent(filepath.Join(dir, "missing.txt")); err == nil {
		t.Error("Expected error for missing file")
	}
}

func TestFileExists(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "f.txt")
	if FileExists(path) {
		t.Error("Expected file to not exist yet")
	}
	if err := WriteFileContent(path, "x"); err != nil {
		t.Fatalf("WriteFileContent failed: %v", err)
	}
	if !FileExists(path) {
		t.Error("Expected file to exist after write")
	}
}

func TestJoinPath(t *testing.T) {
	if got := JoinPath("a", "b", "c.txt"); got != "a/b/c.txt" {
		t.Errorf("Expected \"a/b/c.txt\", got %q", got)
	}
	// 清理冗余分隔符与 "../" 段
	if got := JoinPath("a/", "/b/", "../c"); got != "a/c" {
		t.Errorf("Expected \"a/c\", got %q", got)
	}
}

func TestListFiles(t *testing.T) {
	dir := t.TempDir()
	if err := WriteFileContent(filepath.Join(dir, "a.txt"), "1"); err != nil {
		t.Fatal(err)
	}
	if err := WriteFileContent(filepath.Join(dir, "sub", "b.log"), "2"); err != nil {
		t.Fatal(err)
	}
	if err := WriteFileContent(filepath.Join(dir, "sub", "deep", "c.txt"), "3"); err != nil {
		t.Fatal(err)
	}

	files, err := ListFiles(dir)
	if err != nil {
		t.Fatalf("ListFiles failed: %v", err)
	}
	if len(files) != 3 {
		t.Errorf("Expected 3 files, got %d: %v", len(files), files)
	}

	// root 不存在时返回错误
	if _, err := ListFiles(filepath.Join(dir, "nonexist")); err == nil {
		t.Error("Expected error for missing root")
	}
}

func TestCountFilesByExt(t *testing.T) {
	dir := t.TempDir()
	_ = WriteFileContent(filepath.Join(dir, "a.txt"), "1")
	_ = WriteFileContent(filepath.Join(dir, "b.txt"), "2")
	_ = WriteFileContent(filepath.Join(dir, "c.log"), "3")

	count, err := CountFilesByExt(dir, ".txt")
	if err != nil {
		t.Fatalf("CountFilesByExt failed: %v", err)
	}
	if count != 2 {
		t.Errorf("Expected 2 .txt files, got %d", count)
	}

	count, err = CountFilesByExt(dir, "")
	if err != nil {
		t.Fatalf("CountFilesByExt failed: %v", err)
	}
	if count != 3 {
		t.Errorf("Expected 3 files with empty ext, got %d", count)
	}
}

func TestSafeJoin(t *testing.T) {
	root := t.TempDir()

	got, err := SafeJoin(root, "a/b.txt")
	if err != nil {
		t.Fatalf("SafeJoin failed: %v", err)
	}
	want := filepath.Join(root, "a", "b.txt")
	if got != want {
		t.Errorf("Expected %s, got %s", want, got)
	}

	// 空 sub 返回 root 本身
	got, err = SafeJoin(root, "")
	if err != nil {
		t.Fatalf("SafeJoin with empty sub failed: %v", err)
	}
	if got != root {
		t.Errorf("Expected %s, got %s", root, got)
	}

	// 路径穿越应报错
	for _, evil := range []string{"..", "../escape.txt", "a/../../evil.txt"} {
		if _, err := SafeJoin(root, evil); err == nil {
			t.Errorf("Expected escape error for %q", evil)
		}
	}
}
