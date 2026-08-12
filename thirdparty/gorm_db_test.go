package thirdparty

import (
	"testing"

	"gorm.io/gorm"
)

// newTestDB 创建测试用内存数据库
// 作用：每个测试独立打开一个 :memory: 数据库，测试之间互不影响
func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := OpenSQLiteMemory()
	if err != nil {
		t.Fatalf("Failed to open in-memory sqlite: %v", err)
	}
	return db
}

func TestOpenSQLiteMemory(t *testing.T) {
	db, err := OpenSQLiteMemory()
	if err != nil {
		t.Fatalf("Failed to open in-memory sqlite: %v", err)
	}
	if db == nil {
		t.Error("Expected non-nil db")
	}

	// 验证 AutoMigrate 已建表
	if !db.Migrator().HasTable(&User{}) {
		t.Error("Expected users table exists")
	}
	if !db.Migrator().HasTable(&Order{}) {
		t.Error("Expected orders table exists")
	}
}

func TestUserCRUD(t *testing.T) {
	db := newTestDB(t)

	// Create
	user := &User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	if err := CreateUser(db, user); err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	if user.ID == 0 {
		t.Error("Expected auto-increment ID after create")
	}

	// First 按主键查询
	found, err := FindUserByID(db, user.ID)
	if err != nil {
		t.Fatalf("FindUserByID failed: %v", err)
	}
	if found.Name != "Alice" {
		t.Errorf("Expected name 'Alice', got '%s'", found.Name)
	}

	// Where 条件查询
	users, err := FindUsersByAge(db, 25)
	if err != nil {
		t.Fatalf("FindUsersByAge failed: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Expected 1 user with age 25, got %d", len(users))
	}

	// Update
	if err := UpdateUserAge(db, user.ID, 26); err != nil {
		t.Fatalf("UpdateUserAge failed: %v", err)
	}
	found, _ = FindUserByID(db, user.ID)
	if found.Age != 26 {
		t.Errorf("Expected age 26 after update, got %d", found.Age)
	}

	// Delete
	if err := DeleteUser(db, user.ID); err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
	if _, err := FindUserByID(db, user.ID); err == nil {
		t.Error("Expected error when finding deleted user")
	}
}

func TestFindUsersByAgeNoMatch(t *testing.T) {
	db := newTestDB(t)
	if err := CreateUser(db, &User{Name: "Bob", Age: 30}); err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	users, err := FindUsersByAge(db, 99)
	if err != nil {
		t.Fatalf("FindUsersByAge failed: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(users))
	}
}

func TestFindUserWithOrders(t *testing.T) {
	db := newTestDB(t)

	user := &User{Name: "Carol", Age: 28, Email: "carol@example.com"}
	if err := CreateUser(db, user); err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// 创建关联订单（批量插入）
	orders := []Order{
		{UserID: user.ID, Product: "键盘", Price: 299},
		{UserID: user.ID, Product: "鼠标", Price: 99},
	}
	if err := db.Create(&orders).Error; err != nil {
		t.Fatalf("Create orders failed: %v", err)
	}

	// Preload 预加载关联
	found, err := FindUserWithOrders(db, user.ID)
	if err != nil {
		t.Fatalf("FindUserWithOrders failed: %v", err)
	}
	if len(found.Orders) != 2 {
		t.Errorf("Expected 2 orders, got %d", len(found.Orders))
	}
	if found.Orders[0].Product != "键盘" {
		t.Errorf("Expected product '键盘', got '%s'", found.Orders[0].Product)
	}
}
