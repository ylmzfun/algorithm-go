package stdlib

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestMarshalUser(t *testing.T) {
	u := User{
		ID:        1,
		Name:      "小明",
		Email:     "xiaoming@example.com",
		Age:       18,
		Tags:      []string{"golang", "algorithm"},
		CreatedAt: time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC),
	}
	data, err := MarshalUser(u)
	if err != nil {
		t.Fatalf("MarshalUser failed: %v", err)
	}
	jsonStr := string(data)
	for _, want := range []string{`"id":1`, `"name":"小明"`, `"created_at":"2024-01-15T10:30:00Z"`} {
		if !strings.Contains(jsonStr, want) {
			t.Errorf("Expected JSON to contain %s, got %s", want, jsonStr)
		}
	}
}

func TestMarshalUserOmitEmpty(t *testing.T) {
	// Email 为空、Tags 为 nil 时，omitempty 会省略对应字段
	u := User{ID: 2, Name: "小红", Age: 20}
	data, err := MarshalUser(u)
	if err != nil {
		t.Fatalf("MarshalUser failed: %v", err)
	}
	jsonStr := string(data)
	if strings.Contains(jsonStr, "email") {
		t.Errorf("Expected email omitted, got %s", jsonStr)
	}
	if strings.Contains(jsonStr, "tags") {
		t.Errorf("Expected tags omitted, got %s", jsonStr)
	}
}

func TestUnmarshalUser(t *testing.T) {
	data := []byte(`{"id":1,"name":"小明","age":18,"created_at":"2024-01-15T10:30:00Z"}`)
	u, err := UnmarshalUser(data)
	if err != nil {
		t.Fatalf("UnmarshalUser failed: %v", err)
	}
	if u.ID != 1 || u.Name != "小明" || u.Age != 18 {
		t.Errorf("Expected id/name/age 1/小明/18, got %d/%s/%d", u.ID, u.Name, u.Age)
	}
	// 可选字段缺失时为零值
	if u.Email != "" {
		t.Errorf("Expected empty email, got %s", u.Email)
	}
	if !u.CreatedAt.Equal(time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)) {
		t.Errorf("Expected parsed time, got %v", u.CreatedAt)
	}
}

func TestUnmarshalUserInvalidJSON(t *testing.T) {
	if _, err := UnmarshalUser([]byte(`{invalid`)); err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestProductRoundTrip(t *testing.T) {
	p := Product{Name: "钢笔", Price: 1250} // 12.50 元
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal Product failed: %v", err)
	}
	want := `{"name":"钢笔","price":"12.50"}`
	if string(data) != want {
		t.Errorf("Expected %s, got %s", want, string(data))
	}

	var back Product
	if err := json.Unmarshal(data, &back); err != nil {
		t.Fatalf("Unmarshal Product failed: %v", err)
	}
	if back.Name != "钢笔" || back.Price != 1250 {
		t.Errorf("Expected 钢笔/1250, got %s/%d", back.Name, back.Price)
	}
}

func TestProductInvalidPrice(t *testing.T) {
	var p Product
	if err := json.Unmarshal([]byte(`{"name":"x","price":"abc"}`), &p); err == nil {
		t.Error("Expected error for invalid price")
	}
}

func TestYuan(t *testing.T) {
	p := Product{Name: "书", Price: 3990}
	if got := p.Yuan(); got != 39.9 {
		t.Errorf("Expected 39.9, got %v", got)
	}
}

func TestOrderDateTime(t *testing.T) {
	o := Order{
		OrderID:   1001,
		Product:   "笔记本电脑",
		Amount:    6999.0,
		CreatedAt: DateTime(time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)),
	}
	data, err := json.Marshal(o)
	if err != nil {
		t.Fatalf("Marshal Order failed: %v", err)
	}
	want := `{"order_id":1001,"product":"笔记本电脑","amount":6999,"created_at":"2024-01-15 10:30:00"}`
	if string(data) != want {
		t.Errorf("Expected %s, got %s", want, string(data))
	}

	var back Order
	if err := json.Unmarshal(data, &back); err != nil {
		t.Fatalf("Unmarshal Order failed: %v", err)
	}
	if got := time.Time(back.CreatedAt).Format(CustomLayout); got != "2024-01-15 10:30:00" {
		t.Errorf("Expected custom time, got %s", got)
	}
}

func TestToJSONFromJSON(t *testing.T) {
	u := User{ID: 3, Name: "测试", Age: 25}
	jsonStr, err := ToJSON(u)
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	var back User
	if err := FromJSON(jsonStr, &back); err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}
	if back.Name != "测试" || back.ID != 3 {
		t.Errorf("Expected 测试/3, got %s/%d", back.Name, back.ID)
	}
}

func TestFromJSONInvalid(t *testing.T) {
	var u User
	if err := FromJSON("{broken", &u); err == nil {
		t.Error("Expected error for invalid JSON string")
	}
}

func TestDecodeStrict(t *testing.T) {
	// 未知字段时报错
	if err := DecodeStrict([]byte(`{"id":1,"unknown":true}`), &User{}); err == nil {
		t.Error("Expected error for unknown field")
	}
	// 合法输入正常解析
	u := User{}
	if err := DecodeStrict([]byte(`{"id":1,"name":"n","age":1}`), &u); err != nil {
		t.Errorf("DecodeStrict failed: %v", err)
	}
	if u.ID != 1 {
		t.Errorf("Expected id 1, got %d", u.ID)
	}
}
