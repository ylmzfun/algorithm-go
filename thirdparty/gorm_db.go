package thirdparty

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// User 用户模型（一对多关系中「一」的一方）
// 思路：通过 gorm tag 声明主键、字段约束与外键关联
// 作用：映射 users 表，Orders 字段用于预加载关联订单
// 业务场景：电商系统用户表、内容平台作者表
type User struct {
	ID     uint    `gorm:"primaryKey"`       // 主键
	Name   string  `gorm:"size:64;not null"` // 用户名
	Age    int     // 年龄
	Email  string  // 邮箱
	Orders []Order `gorm:"constraint:OnDelete:CASCADE"` // 关联订单列表
}

// Order 订单模型（一对多关系中「多」的一方）
// 思路：通过 UserID 外键关联 User，主键自增
// 作用：映射 orders 表，是 Preload 预加载的目标表
type Order struct {
	ID      uint    `gorm:"primaryKey"` // 主键
	UserID  uint    `gorm:"index"`      // 外键，自动创建索引
	Product string  // 商品名称
	Price   float64 // 商品价格
}

// OpenSQLiteMemory 打开内存 SQLite 并自动建表
// 思路：使用纯 Go 驱动 glebarez/sqlite 打开 :memory: 数据库，
// 再执行 AutoMigrate 建表；限制单连接保证内存库数据一致
// 作用：无需 CGO 即可获得一个可用的 ORM 数据库实例，测试结束后自动销毁
// 业务场景：单元测试、原型演示等需要临时数据库的场景
func OpenSQLiteMemory() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	// 内存数据库需限制为单一连接，否则不同连接各自持有独立的内存库
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取底层连接失败: %w", err)
	}
	sqlDB.SetMaxOpenConns(1)

	if err := db.AutoMigrate(&User{}, &Order{}); err != nil {
		return nil, fmt.Errorf("自动建表失败: %w", err)
	}
	return db, nil
}

// CreateUser 创建用户
// 作用：向 users 表插入一条记录，回填自增主键
// 复杂度：O(1)
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// FindUserByID 按主键查询用户
// 作用：演示 First 主键查询，记录不存在时返回 gorm.ErrRecordNotFound
// 复杂度：O(1)（主键索引）
func FindUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUsersByAge 按年龄条件查询
// 作用：演示 Where 条件查询，返回满足条件的所有用户
// 复杂度：O(n)
func FindUsersByAge(db *gorm.DB, age int) ([]User, error) {
	var users []User
	if err := db.Where("age = ?", age).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUserAge 更新用户年龄
// 作用：演示 Update 单字段更新
// 复杂度：O(1)（主键定位）
func UpdateUserAge(db *gorm.DB, id uint, age int) error {
	return db.Model(&User{}).Where("id = ?", id).Update("age", age).Error
}

// DeleteUser 按主键删除用户
// 作用：演示 Delete 删除，关联订单由外键级联约束一并删除
// 复杂度：O(1)（主键定位）
func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
}

// FindUserWithOrders 预加载关联订单查询
// 作用：演示 Preload 一次查出用户及其关联订单，避免 N+1 查询问题
// 复杂度：固定 2 条 SQL（主表 + 关联表）
func FindUserWithOrders(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.Preload("Orders").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
