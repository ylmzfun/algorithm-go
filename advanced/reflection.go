package advanced

import (
	"errors"
	"fmt"
	"reflect"
)

// reflect 包提供运行时反射能力，可以在运行时检查类型、读取值、动态调用函数与方法：
// 1. reflect.TypeOf(v)：获取静态类型信息；reflect.ValueOf(v)：获取值信息
// 2. 结构体反射：遍历字段、读取 struct tag（json、validate 等）
// 3. 动态调用：通过 reflect.Value.Call 在运行时调用函数与方法
// 4. 典型用途：序列化框架（json/yaml）、ORM、测试工具、依赖注入
// 注意：反射有运行时开销且绕过静态类型检查，业务代码应谨慎使用

// User 用于反射演示的结构体（含 struct tag）
// 业务场景：结构体 tag 常见于 JSON 序列化、表单校验、ORM 字段映射
type User struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0,lte=150"`
}

// TypeName 返回值的动态类型名称
// 思路：reflect.TypeOf 获取类型信息，取其 Name()
// 作用：演示最基本的类型反射
// 业务场景：日志中打印对象类型、通用处理器的类型分发
func TypeName(v interface{}) string {
	rt := reflect.TypeOf(v)
	if rt == nil {
		return "<nil>" // 反射 nil 接口
	}
	return rt.Name()
}

// KindName 返回值的底层类型类别（Kind）
// 作用：区分 Kind 与具体类型——TypeName 得到 User，KindName 得到 struct
// 业务场景：根据底层类别（struct/slice/map/ptr）走不同的通用处理逻辑
func KindName(v interface{}) string {
	return reflect.TypeOf(v).Kind().String()
}

// FieldInfo 结构体字段的反射信息
type FieldInfo struct {
	Name  string      // 字段名
	Type  string      // 字段类型名
	Tag   string      // 原始 tag 字符串
	Value interface{} // 字段值
}

// FieldValues 遍历结构体的所有字段，返回每个字段的名称、类型、tag 与值
// 思路：reflect.ValueOf(v).Type().Field(i) 取字段元信息，Field(i).Interface() 取字段值
// 作用：演示结构体字段遍历与 struct tag 读取
// 业务场景：通用校验器、对象转 map、动态表单生成
func FieldValues(v interface{}) ([]FieldInfo, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("FieldValues requires a struct")
	}

	rt := rv.Type()
	infos := make([]FieldInfo, 0, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		infos = append(infos, FieldInfo{
			Name:  field.Name,
			Type:  field.Type.Name(),
			Tag:   string(field.Tag),
			Value: rv.Field(i).Interface(),
		})
	}
	return infos, nil
}

// GetTagValue 读取结构体指定字段的某个 tag 值
// 思路：FieldByName 查找字段，Tag.Get(tagKey) 读取对应 tag
// 作用：演示 struct tag 的按需读取
// 业务场景：JSON 序列化器读取 json tag、校验框架读取 validate tag
// 错误：输入非结构体或字段不存在时返回错误；字段存在但无该 tag 时返回空字符串
func GetTagValue(v interface{}, fieldName, tagKey string) (string, error) {
	rt := reflect.TypeOf(v)
	if rt.Kind() != reflect.Struct {
		return "", errors.New("GetTagValue requires a struct")
	}
	field, ok := rt.FieldByName(fieldName)
	if !ok {
		return "", fmt.Errorf("field %q not found", fieldName)
	}
	return field.Tag.Get(tagKey), nil
}

// CallFunction 通过反射动态调用函数
// 思路：将参数转为 []reflect.Value，Call 调用后把结果转回 []interface{}
// 作用：演示基于反射的函数调用（函数在运行时才确定）
// 业务场景：插件系统、注册表模式、根据配置动态选择处理函数
// 错误：fn 不是函数、参数个数不匹配、参数类型不匹配时返回错误
func CallFunction(fn interface{}, args ...interface{}) (results []interface{}, err error) {
	fv := reflect.ValueOf(fn)
	if fv.Kind() != reflect.Func {
		return nil, errors.New("CallFunction requires a function")
	}
	if len(args) != fv.Type().NumIn() {
		return nil, fmt.Errorf("argument count mismatch: got %d, want %d",
			len(args), fv.Type().NumIn())
	}

	// 参数类型不匹配时 reflect 会 panic，这里用 defer recover 转成 error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("CallFunction failed: %v", r)
			results = nil
		}
	}()

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	out := fv.Call(in)
	results = make([]interface{}, len(out))
	for i, rv := range out {
		results[i] = rv.Interface()
	}
	return results, nil
}

// Calculator 用于反射动态方法调用演示的结构体
type Calculator struct {
	base int
}

// NewCalculator 创建 Calculator
func NewCalculator(base int) *Calculator {
	return &Calculator{base: base}
}

// Add 将 base 与参数相加（导出方法才能被 MethodByName 找到）
func (c *Calculator) Add(n int) int {
	return c.base + n
}

// CallMethod 通过反射动态调用结构体的方法
// 思路：MethodByName 查找方法，再复用 CallFunction 完成调用
// 作用：演示基于反射的方法调用
// 业务场景：路由分发（根据请求动态调用不同 Handler）、测试框架
// 错误：方法不存在、参数不匹配时返回错误
func CallMethod(v interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	mv := reflect.ValueOf(v).MethodByName(methodName)
	if !mv.IsValid() {
		return nil, fmt.Errorf("method %q not found", methodName)
	}
	return CallFunction(mv.Interface(), args...)
}
