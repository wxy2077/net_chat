package pkg

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"log"
	"os"
)

var PwdJwt = &pwdJwt{}

type pwdJwt struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func (p *pwdJwt) LoadConfig(privateKey, pubKeyPath string) {
	publicKeyByte, err := os.ReadFile(privateKey)
	if err != nil {
		log.Fatalf("fail to read file jwt public key, err: %s", err.Error())
	}
	p.publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)

	privateKeyByte, err := os.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("fail to read file jwt private key, err: %s", err.Error())
	}
	p.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		log.Fatalf("fail to parse rsa private key from pem, err: %s", err.Error())
	}
}

func (p *pwdJwt) GenerateJwtToken(claims jwt.MapClaims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err = token.SignedString(p.privateKey)
	return
}

func (p *pwdJwt) ParseToken(tokenStr string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("验证Token的加密类型错误")
		}
		return p.publicKey, nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Token无效或者无对应值")

}
