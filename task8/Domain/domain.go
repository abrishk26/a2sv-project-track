package domain

import "context"

type Task struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	DueDate     string `json:"due_date" bson:"due_date"`
	Done        bool   `json:"done" bson:"done"`
}

type User struct {
	ID           string `json:"id" bson:"_id"`
	Username     string `json:"username" bson:"username"`
	IsAdmin      bool   `json:"is_admin" bson:"is_admin"`
	PasswordHash string `json:"-" bson:"password_hash"`
}

type IUserRepository interface {
	Add(ctx context.Context, u User) error
	Get(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, u User) error
	GetAll(ctx context.Context) (*[]User, error)
}

type ITaskRepository interface {
	Add(ctx context.Context, t Task) error
	Get(ctx context.Context, id string) (*Task, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, t Task) error
	GetAll(ctx context.Context) (*[]Task, error)
}

type IPasswordService interface {
	Hash(password string) (string, error)
	Verify(password, hash string) error
}

type ITokenService interface {
	GenerateToken(userID string) (string, error)
	VerifyToken(token string) (string, error)
}
