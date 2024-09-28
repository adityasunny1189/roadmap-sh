package dtos

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AuthResponse struct {
	BearerToken string `json:"BearerToken"`
}
