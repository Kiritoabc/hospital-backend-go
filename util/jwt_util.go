package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

type MyCustomClaims struct {
	Phone string
	jwt.RegisteredClaims
}

// 签名密钥
const signKey = "hello jwt"

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(strLen int) string {
	randBytes := make([]rune, strLen)
	for i := range randBytes {
		randBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randBytes)
}

func GenerateTokenUsingHs256(phone string) (string, error) {
	claim := MyCustomClaims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "牵风散步的雲",                                        // 签发者
			Subject:   "User",                                          // 签发对象
			Audience:  jwt.ClaimStrings{"hospital", "web"},             //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(signKey))
	return token, err
}

func ParseTokenHs256(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
