package auth

type AuthoriztaionQuery struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type usersRecord struct {
	Role     Role   `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}
