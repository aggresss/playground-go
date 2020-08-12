// user.go
package user

// User 表示一个用户
type User struct {
	Name string
}

// UserRepository 用户仓库
type UserRepository interface {
	// 根据用户id查询得到一个用户或是错误信息
	FindOne(id int) (*User, error)
}
