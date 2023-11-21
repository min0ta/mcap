package auth

type AuthoriztaionQuery struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type usersRecord struct {
	Role     Role
	Username string
	Password string
}
