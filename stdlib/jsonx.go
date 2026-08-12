package stdlib

// 思路：encoding/json 基于反射在 Go 结构体与 JSON 文本之间互转，通过结构体 tag
// 定制字段名、可选性，通过自定义 MarshalJSON/UnmarshalJSON 扩展编解码逻辑。
// 作用：提供结构化的数据交换能力。
// 业务场景：
// 1. REST API 的请求/响应体
// 2. 配置文件与缓存数据的序列化
// 3. 消息队列的消息体编解码

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

// User 用户信息，演示字段 tag 与 omitempty
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty"` // 为空时省略该字段
	Age       int       `json:"age"`
	Tags      []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at"` // time.Time 默认按 RFC3339 输出
}

// MarshalUser 将 User 序列化为 JSON 字节
func MarshalUser(u User) ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalUser 将 JSON 字节解析为 User
// 未知字段会被忽略；可选字段缺失时保持零值
func UnmarshalUser(data []byte) (User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	return u, err
}

// Product 商品，演示自定义序列化
// 内部价格以"分"（int64）存储，避免浮点误差；对外 JSON 以"元"（字符串）展示
type Product struct {
	Name  string `json:"name"`
	Price int64  `json:"-"` // 不直接参与默认序列化，由 MarshalJSON 接管
}

// MarshalJSON 自定义序列化：将价格由分转换为元（保留两位小数）
func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name  string `json:"name"`
		Price string `json:"price"` // 单位：元
	}{
		Name:  p.Name,
		Price: fmt.Sprintf("%.2f", float64(p.Price)/100.0),
	})
}

// UnmarshalJSON 自定义反序列化：将价格由元转换回分
// 价格格式非法时返回错误
func (p *Product) UnmarshalJSON(data []byte) error {
	var raw struct {
		Name  string `json:"name"`
		Price string `json:"price"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	yuan, err := strconv.ParseFloat(raw.Price, 64)
	if err != nil {
		return fmt.Errorf("invalid price %q: %w", raw.Price, err)
	}
	p.Name = raw.Name
	p.Price = int64(math.Round(yuan * 100))
	return nil
}

// Yuan 返回以元为单位的价格（仅用于展示，金额计算仍应使用 Price）
func (p Product) Yuan() float64 {
	return float64(p.Price) / 100.0
}

// DateTime 自定义时间类型，JSON 中使用 "2006-01-02 15:04:05" 布局
type DateTime time.Time

// MarshalJSON 按自定义布局序列化
func (dt DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(dt).Format(CustomLayout))
}

// UnmarshalJSON 按自定义布局反序列化，格式非法时返回错误
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(CustomLayout, s)
	if err != nil {
		return err
	}
	*dt = DateTime(t)
	return nil
}

// Order 订单，演示自定义时间类型的应用
type Order struct {
	OrderID   int64    `json:"order_id"`
	Product   string   `json:"product"`
	Amount    float64  `json:"amount"`
	CreatedAt DateTime `json:"created_at"`
}

// ToJSON 将任意值序列化为 JSON 字符串
func ToJSON(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

// FromJSON 将 JSON 字符串反序列化到 v
// v 必须为指针；解析失败时返回错误
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

// DecodeStrict 严格模式解析：JSON 中出现结构体未定义的字段时返回错误
func DecodeStrict(data []byte, v interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}
