package services

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService(email string, password string) LoginService {
	return &loginInformation{
		email:    email,
		password: password,
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
