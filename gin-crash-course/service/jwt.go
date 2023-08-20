package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// custom data to be embedded in JWT token
type jwtCustomClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "mateuszmidor.com",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secretkey"
	}
	return secret
}

func (s *jwtService) GenerateToken(name string, admin bool) string {
	claims := &jwtCustomClaims{
		name,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// ValidateToken is simply : check if provided tokenString is signed with our secret key
// and if so - parse and return the token
func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	tryGetSecretKey := func(token *jwt.Token) (interface{}, error) {
		// check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	}

	return jwt.Parse(tokenString, tryGetSecretKey)
}
