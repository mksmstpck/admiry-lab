package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"

	"github.com/pborman/uuid"
)

type JWTs interface {
	Create(ttl time.Duration, secret_key []byte, userID uuid.UUID) (string, error)
	Validate(token string, public_key []byte) (uuid.UUID, error)
}

type JWT struct {
	secret_key []byte
	public_key []byte
}

func NewJWT(secret_key []byte, public_key []byte) *JWT {
	return &JWT{
		secret_key: secret_key,
		public_key: public_key,
	}
}

func (j *JWT) Create(ttl time.Duration, secret_key []byte, userID uuid.UUID) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(secret_key)

	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID.String()
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		log.Error("utils.Create: ", err)
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) Validate(token string, public_key []byte) (uuid.UUID, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.public_key)
	if err != nil {
		log.Error("utils.Validate: ", err)
		return nil, err
	}

	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Error("utils.Validate: ", err)
			return nil, errors.New("unexpected signing method")
		}
		return key, nil
	})
	if err != nil {
		log.Error("utils.Validate: ", err)
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		log.Error("utils.Validate: ", err)
		return nil, errors.New("invalid claims")
	}

	return claims["user_id"].(uuid.UUID), nil
}
