package token

import (
	"io/ioutil"
	"time"

	setting "github.com/CircleYu/GameServer/setting"
	"github.com/dgrijalva/jwt-go"
)

var hmacSecret []byte

type myCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func init() {
	// Load sample key data
	if keyData, e := ioutil.ReadFile(setting.HmacKeyPath()); e == nil {
		hmacSecret = keyData
	} else {
		panic(e)
	}
}

func CreateToken(userName string) (string, error) {

	claims := myCustomClaims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(15)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(hmacSecret)
}
