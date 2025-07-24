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
