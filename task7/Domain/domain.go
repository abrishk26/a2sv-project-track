package domain

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
	Role         string `json:"role" bson:"role"`
	PasswordHash string `json:"-" bson:"password_hash"`
}

type IUserRepository interface {
	Add(u User) error
	Get(id string) (*User, error)
	Delete(id string) error
	Update(id string, u User) error
	GetAll() (*[]User, error)
}

type ITaskRepository interface {
	Add(t Task) error
	Get(id string) (*Task, error)
	Delete(id string) error
	Update(id string, t Task) error
	GetAll() (*[]Task, error)
}
