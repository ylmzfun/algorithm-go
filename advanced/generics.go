package advanced

// 泛型（Generics，Go 1.18+）允许编写与类型无关的代码：
// 1. 类型参数 [T constraint]：函数或类型可声明一个或多个类型参数
// 2. 约束（constraint）：限制类型参数可接受的具体类型，可以是接口或类型集合
//    - comparable：内置约束，表示可用 ==、!= 比较的类型
//    - any：等价于 interface{}，无任何约束
//    - 自定义约束：如下方 ordered，用类型集合语法 ~int | ~string | ...
// 3. 相比 interface{} + 类型断言，泛型在编译期确定类型，更安全且无运行时开销

// ordered 自定义约束：所有支持 <、<=、>、>= 比较的内置有序类型
type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Max 泛型最大值函数
// 约束：ordered（自定义有序类型约束）
// 作用：一个函数同时支持 int、float64、string 等所有有序类型
// 业务场景：通用工具库中的比较辅助函数
func Max[T ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min 泛型最小值函数
// 约束：ordered
// 作用：与 Max 对称的最小值实现
// 业务场景：通用工具库中的比较辅助函数
func Min[T ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Filter 泛型过滤函数：保留满足谓词条件的元素
// 约束：any（不限制元素类型，仅需提供判断函数）
// 作用：一个实现适用于任意元素类型
// 业务场景：数据清洗（过滤空值）、批量筛选
// 时间复杂度：O(n)
func Filter[T any](slice []T, keep func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map 泛型映射函数：对每个元素应用变换函数，返回新切片
// 类型参数：T 为输入元素类型，U 为输出元素类型（均为任意类型）
// 作用：演示多类型参数的泛型函数
// 业务场景：批量转换（ID 列表转对象、枚举转字符串）
// 时间复杂度：O(n)
func Map[T, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

// Contains 泛型包含判断：切片中是否存在目标元素
// 约束：comparable——只有可比较的类型才能用 == 判断相等
// 作用：演示 comparable 约束的用途
// 业务场景：白名单校验、去重前置检查
// 时间复杂度：O(n)
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// Stack 泛型栈容器
// 思路：用切片作为底层存储，Push 在末尾追加、Pop 取末尾元素
// 作用：演示泛型容器——同一份实现可容纳任意类型
// 业务场景：括号匹配、表达式求值、撤销操作历史
type Stack[T any] struct {
	data []T
}

// NewStack 创建空的泛型栈
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push 将元素压入栈顶
// 时间复杂度：平均 O(1)
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop 弹出栈顶元素；栈为空时返回零值和 false
// 时间复杂度：O(1)
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

// Peek 查看栈顶元素但不弹出；栈为空时返回零值和 false
// 时间复杂度：O(1)
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Size 返回栈中元素个数
// 时间复杂度：O(1)
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// IsEmpty 判断栈是否为空
// 时间复杂度：O(1)
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
