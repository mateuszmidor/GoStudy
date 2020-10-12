package service

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	user, password string
}

func NewLoginService() LoginService {
	return &loginService{
		user:     "admin",
		password: "pass"}
}

func (s *loginService) Login(username, password string) bool {
	return username == s.user && password == s.password
}
