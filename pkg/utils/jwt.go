package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

func CreateRefreshToken(claims jwt.MapClaims) string {
	privateKey := ConvertToPrivateKey(viper.GetString("refresh_token_private_key"))
	exp := viper.GetInt("refresh_token_exp")

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(exp)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	CatchError(err)

	return tokenString
}

func CreateAccessToken(claims jwt.MapClaims) string {
	privateKey := ConvertToPrivateKey(viper.GetString("access_token_private_key"))
	exp := viper.GetInt("access_token_exp")

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(exp)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	CatchError(err)

	return tokenString
}

func VerifyRefreshToken(tokenString string) (*jwt.MapClaims, error) {
	publicKey := ConvertToPublicKey(viper.GetString("refresh_token_public_key"))
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}

func VerifyAccessToken(tokenString string) (*jwt.MapClaims, error) {
	publicKey := ConvertToPublicKey(viper.GetString("access_token_public_key"))
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}

func ExtractJwtExp(tokenString string) time.Time {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		CatchError(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		CatchError(fmt.Errorf("exp claim not found or not valid"))
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		CatchError(fmt.Errorf("invalid token claims"))
	}

	return time.Unix(int64(exp), 0)
}
