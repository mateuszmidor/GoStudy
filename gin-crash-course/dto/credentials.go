package dto

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

func NewJWT(token string) *JWT {
	return &JWT{Token: token}
}
