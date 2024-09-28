package domain

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	FullName     string `json:"full_name"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}
