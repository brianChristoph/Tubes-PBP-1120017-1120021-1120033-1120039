package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	m "github.com/Tubes-PBP/models"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(id int, name string, password string, email string, userType string, balance int) string
	ValidateToken(token string) (*jwt.Token, error)
	ValidateTokenFromCookies(r *http.Request) (bool, m.User)
}
type authCustomClaims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	Balance  int    `json:"balance"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService(name string) JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    name,
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(id int, name string, password string, email string, userType string, balance int) string {
	claims := &authCustomClaims{
		id,
		name,
		password,
		email,
		userType,
		balance,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func ResetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "TOKEN",
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func (service *jwtServices) ValidateTokenFromCookies(r *http.Request) (bool, m.User) {
	var user m.User
	if cookie, err := r.Cookie("TOKEN"); err == nil {
		accessToken := cookie.Value
		accessClaims := &authCustomClaims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return []byte(service.secretKey), nil
		})
		if err == nil && parsedToken.Valid {
			user.ID = accessClaims.ID
			user.Name = accessClaims.Name
			user.Password = accessClaims.Password
			user.Email = accessClaims.Email
			user.UserType = accessClaims.UserType
			user.Balance = accessClaims.Balance
			return true, user
		}
	}
	return false, user
}
