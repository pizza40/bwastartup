package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
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