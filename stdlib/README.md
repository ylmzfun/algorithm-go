# Go 标准库介绍与用法

Go 标准库是 Go 语言自带的官方库，覆盖输入输出、网络、序列化、并发、反射等日常开发所需的大部分能力，无需任何第三方依赖。本目录提供 7 组可直接运行的示例代码（`http.go`、`jsonx.go`、`iox.go`、`filex.go`、`regexpx.go`、`timex.go`、`sortx.go`），每组均配套单元测试。

本文按功能分类介绍常用标准库包：用途说明 + 典型用法 + 本项目的示例位置。

## 目录结构

| 文件 | 对应标准库包 | 演示内容 |
|---|---|---|
| `http.go` | `net/http`、`net/http/httptest` | HTTP 服务端路由、JSON 接口、客户端请求 |
| `jsonx.go` | `encoding/json` | 序列化/反序列化、字段 tag、自定义编解码 |
| `iox.go` | `io`、`bufio` | 按行读取、缓冲写入、单词统计 |
| `filex.go` | `os`、`path/filepath` | 文件读写、目录遍历、防路径穿越 |
| `regexpx.go` | `regexp` | 邮箱/手机号校验、分组提取、脱敏 |
| `timex.go` | `time` | 格式化、解析、定时器、年龄计算 |
| `sortx.go` | `sort`、`slices` | 自定义排序、二分查找 |

> 并发相关（`sync`、`context`）、反射（`reflect`）等标准库的用法见 `advanced/` 目录。

---

## 1. 输入输出：fmt / io / bufio / os / strconv / strings

### fmt — 格式化输入输出
最常用的格式化输出库，`Printf` 支持丰富的格式动词。

```go
fmt.Printf("name=%s age=%d ok=%t\n", "Alice", 25, true) // name=Alice age=25 ok=true
s := fmt.Sprintf("%.2f", 3.14159)                        // 保留两位小数: "3.14"
fmt.Errorf("连接失败: %w", err)                            // %w 包装错误，配合 errors.Is/As 使用
```

### io — 读写抽象
`io.Reader` / `io.Writer` 是 Go 流式 I/O 的基石接口，网络、文件、内存都实现了它们。

```go
data, _ := io.ReadAll(r)          // 读取所有内容
n, _ := io.Copy(dst, src)         // 拷贝流，返回拷贝字节数
```

### bufio — 缓冲读写
减少系统调用次数，大幅提升读写性能；`Scanner` 可方便地按行分割。

```go
scanner := bufio.NewScanner(reader)
scanner.Split(bufio.ScanLines)    // 按行分割（默认）
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

**本项目示例**：`iox.go` — `NewLineReader` 封装 Scanner 按行读取；`CopyToWriter` 封装 io.Copy；`WriteLines` 批量写入。

### os — 操作系统接口
文件操作、环境变量、进程管理、路径常量。

```go
f, _ := os.Create("tmp.txt")       // 创建文件
os.WriteFile("tmp.txt", []byte("hi"), 0o644)
env := os.Getenv("GOPATH")         // 读取环境变量
os.TempDir()                       // 系统临时目录
```

### strconv — 字符串与基础类型转换
```go
n, _ := strconv.Atoi("42")         // 字符串 -> int
s := strconv.Itoa(42)              // int -> 字符串
f, _ := strconv.ParseFloat("3.14", 64)
```

### strings — 字符串处理
```go
parts := strings.Split("a,b,c", ",")        // 分割
joined := strings.Join(parts, "-")          // 拼接
strings.TrimSpace("  hi  ")                 // 去空白
strings.Contains("hello", "ell")            // 包含
strings.ReplaceAll("a-b-c", "-", "_")       // 全局替换
var b strings.Builder                       // 高效字符串拼接
b.WriteString("hello")
```

---

## 2. 数据结构与算法：sort / slices

### sort — 排序与搜索
`slices`（Go 1.21+）是对切片操作的泛型库，与 `sort` 互为补充。

```go
sort.Ints(nums)                               // 原地升序排序
sort.Slice(people, func(i, j int) bool {      // 自定义排序
    return people[i].Age < people[j].Age
})
idx := sort.Search(n, func(i int) bool { ... }) // 二分查找
```

**本项目示例**：`sortx.go` — `SortByAgeAsc`（sort.Slice）、`SortByNameLen`（slices.SortFunc）、`StableSortByAge`（sort.SliceStable 稳定排序）、`SearchInt` / `SearchInsertPosition`（sort.Search 二分）。

---

## 3. 时间：time

Go 时间格式化的布局是固定参照时间 `2006-01-02 15:04:05`（1 月 2 日 3 点 4 分 5 秒，2006 年），不是 `yyyy-MM-dd` 之类的占位符。

```go
now := time.Now()
now.Format("2006-01-02 15:04:05")             // 自定义格式
now.Format(time.RFC3339)                      // "2026-08-12T14:31:09+08:00"
t, _ := time.Parse(time.RFC3339, "2026-08-12T14:31:09+08:00")

ticker := time.NewTicker(time.Second)         // 周期定时器
for range ticker.C { /* 每秒执行 */ }
<-time.After(2 * time.Second)                 // 一次性延迟
```

**本项目示例**：`timex.go` — `FormatRFC3339` / `ParseRFC3339`、`FormatCustom` / `ParseCustom`、`RunTicker`（返回可安全重复调用的 stop 函数）、`WaitTimeout`、`DayRange`（当天 0 点到 24 点）、`CalculateAge`（年龄计算，处理生日边界）。

---

## 4. 字符串与正则：regexp

正则表达式引擎（RE2 语法），用于文本校验、提取、替换。

```go
re, err := regexp.Compile(`^\d{3}-\d{4}$`)    // 编译失败返回 error，而非 panic
re.MatchString("123-4567")                    // 匹配判断
re.FindAllString(text, -1)                    // 提取所有匹配
re.ReplaceAllString("13812345678", "***")     // 替换
```

**本项目示例**：`regexpx.go` — `IsValidEmail` / `IsValidPhone`（常用格式校验）、`ExtractEmails` / `ExtractDateParts`（分组提取）、`MaskPhone`（`138****5678` 脱敏）、`CompilePattern`（编译并返回错误）。

---

## 5. 网络：net/http / net/http/httptest

标准库自带的 HTTP 框架，足以支撑中小型服务；`httptest` 提供不占用真实端口的测试利器。

```go
// 服务端
mux := http.NewServeMux()
mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"msg": "hello"})
})
http.ListenAndServe(":8080", mux)

// 客户端
resp, _ := http.Get("http://localhost:8080/api/hello")
client := &http.Client{Timeout: 5 * time.Second} // 务必设置超时，防止悬挂
```

**本项目示例**：`http.go` — `NewMux`（ServeMux 路由）、`JSONResponse` / `JSONError`（统一 JSON 响应）、`NewClient` / `GetJSON` / `PostJSON`（带超时的客户端封装）。测试全部基于 `httptest.NewServer` / `httptest.NewRecorder`，无需真实端口。

---

## 6. 序列化：encoding/json

```go
type User struct {
    Name string `json:"name"`             // 字段 tag 映射 JSON 键名
    Age  int    `json:"age,omitempty"`    // omitempty: 零值不输出
}
data, _ := json.Marshal(user)             // 结构体 -> JSON
var u User
json.Unmarshal(data, &u)                  // JSON -> 结构体
```

高级用法：
- 实现 `MarshalJSON` / `UnmarshalJSON` 自定义编解码（如金额分↔元转换、时间格式定制）；
- `json.NewDecoder(r).DisallowUnknownFields()` 严格解码，拒绝未知字段。

**本项目示例**：`jsonx.go` — `MarshalUser` / `UnmarshalUser`、`Product`（自定义 MarshalJSON 将"分"转为"元"字符串）、`DateTime`（自定义时间布局）、`ToJSON` / `FromJSON`（字符串便捷封装）、`DecodeStrict`（严格解码）。

---

## 7. 并发：sync / context / atomic

> 完整示例见 `advanced/` 目录，这里列出常用 API。

### sync — 同步原语
```go
var mu sync.Mutex                          // 互斥锁，保护共享数据
mu.Lock(); defer mu.Unlock()

var rw sync.RWMutex                        // 读写锁：读多写少场景
var once sync.Once                         // 保证初始化只执行一次
var wg sync.WaitGroup                      // 等待一组 goroutine 完成
pool := sync.Pool{New: func() interface{} { ... }} // 对象池，减少内存分配
```

**本项目示例**：`advanced/syncx.go` — `Counter`（Mutex）、`RWCache`（RWMutex）、`GetInstance`（Once 单例）、`ObjectPool`（Pool）。

### context — 上下文
跨 API 边界的取消信号、超时控制、请求级值传递，**作为函数第一个参数**传递。

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
ctx2 := context.WithValue(ctx, key, value) // 传值（key 应为自定义类型，避免冲突）
```

**本项目示例**：`advanced/context.go` — `RunWithTimeout`、`ProcessItems`（协作式取消）、`WithRequestID` / `GetRequestID`（值传递）。

### atomic — 原子操作
无锁的并发安全计数，性能优于 Mutex。

```go
var count atomic.Int64
count.Add(1)
n := count.Load()
```

---

## 8. 反射：reflect

运行时动态检查类型与值：遍历结构体字段、读取 struct tag、动态调用函数与方法。常用于序列化框架、ORM、依赖注入、通用校验器。

```go
t := reflect.TypeOf(v)          // 类型信息
v := reflect.ValueOf(v)         // 值信息
field := t.Field(i)             // 字段元信息（含 tag）
field.Tag.Get("json")           // 读取 tag
v.MethodByName("Add").Call(...) // 动态调用方法
```

> 注意：反射有运行时开销，且绕过编译期类型检查，业务热路径应避免使用。

**本项目示例**：`advanced/reflection.go` — `TypeName` / `KindName`、`FieldValues`（字段遍历）、`GetTagValue`（tag 读取）、`CallFunction` / `CallMethod`（动态调用）。

---

## 9. 错误处理：errors / fmt

Go 用显式的 error 返回值而非异常。Go 1.13+ 支持错误链（wrapping）。

```go
if err != nil {
    return fmt.Errorf("处理用户失败: %w", err)  // %w 包装底层错误
}
// 调用方
if errors.Is(err, sql.ErrNoRows) { ... }    // 链上是否包含某错误
var target *MyError
if errors.As(err, &target) { ... }          // 链上是否有指定类型
```

---

## 10. 其他常用包速览

| 包 | 用途 | 典型用法 |
|---|---|---|
| `math` | 数学函数与常量 | `math.MaxInt64`、`math.Sqrt` |
| `encoding/base64` | 二进制编码 | `base64.StdEncoding.EncodeToString` |
| `crypto/sha256` | 哈希 | `sha256.Sum256(data)` |
| `flag` | 命令行 flag 解析 | `flag.String("name", "", "说明")` |
| `testing` | 单元测试/基准测试 | `TestXxx`、`BenchmarkXxx`、表驱动测试 |
| `log` | 简单日志 | `log.Printf`、`log.SetFlags`（生产环境建议用 `thirdparty/zap_log.go`） |
| `container/heap` | 堆实现 | 优先级队列（见 `heap/` 目录） |
| `unicode/utf8` | UTF-8 编解码 | `utf8.RuneCountInString` |

---

## 运行方式

```bash
# 运行本目录全部测试
go test ./stdlib/...

# 查看标准库应用的演示输出（第 18 节）
go run main.go
```

## 参考

- 官方文档：<https://pkg.go.dev/std>
- 中文翻译文档：<https://studygolang.com/pkgdoc>
