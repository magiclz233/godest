package domain

// Repository 定义用户聚合的持久化接口
type Repository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
	GetAll() ([]User, error)
}
