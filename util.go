package main

import (
	"errors"
	"os"
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

func parsePassword(pwdStr string) (string, error) {
	pwdArr := strings.Split(pwdStr, ":")
	if len(pwdArr) < 1 {
		return "", errors.New("password format error")
	}
	if pwdArr[0] == "1" && len(pwdArr) == 2 {
		return pwdArr[1], nil
	}

	return "", errors.New("password format unknown")
}

type jwtClaims struct {
	jwtgo.StandardClaims
	Info interface{} `json:"info"`
}

func signature(info interface{}, expiresSpan int) (string, error) {
	var jClaims jwtClaims
	jClaims.ExpiresAt = time.Now().Unix() + int64(expiresSpan)
	jClaims.Info = info
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jClaims)
	return token.SignedString([]byte(jwtSecret))
}

func getEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
