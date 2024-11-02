package utils

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
)

func ConvertToPrivateKey(key string) (privateKey *rsa.PrivateKey) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	CatchError(err)

	return privateKey
}

func ConvertToPublicKey(key string) (publicKey *rsa.PublicKey) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	CatchError(err)

	return publicKey
}
