package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BWASTARTUP_sEc12eT_k3Y")

func NewService() *jwtService{
	return &jwtService{}
}

func (s *jwtService) GenerateToken(UserID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = UserID

	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//verify signature
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil{
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error){
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil{
		return token, err
	}

	return token, nil
}