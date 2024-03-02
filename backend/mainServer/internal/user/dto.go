package user

type UserSignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserSignUp struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}
