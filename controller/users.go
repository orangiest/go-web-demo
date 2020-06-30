package controller

import (
	"cwm.wiki/web/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type userStdClaims struct {
	 jwt.StandardClaims
	 *models.User
}

func LoginController(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err!=nil {
		return
	}

	if user.Username == "admin" && user.Password == "admin" {
		token,err := JwtGenerateToken(&user,time.Minute * 20)

		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"err":err})

			return
		}

		c.JSON(http.StatusOK,gin.H{"token":token})
	}

}

func (u userStdClaims) Valid() (err error) {

	if u.VerifyExpiresAt(time.Now().Unix(),true) == false {
		return  errors.New("token is expired")
	}

	return nil

}

func JwtGenerateToken(user *models.User,d time.Duration) (string,error) {
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt: time.Now().Unix(),
	}

	uClaims := userStdClaims{
		stdClaims,
		user,
	}

	// 注意加密方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenString, err := token.SignedString([]byte("Test"))

	if err != nil {
		fmt.Println("It's error")
		return "", err
	}

	return tokenString,nil


}


func JwtParseUser(tokenString string) (*models.User, error) {

	if tokenString == "" {
		return nil,errors.New("no token is found in Authorization Bearer")
	}

	claims := userStdClaims{}

	_, err := jwt.ParseWithClaims(tokenString,&claims,func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("Test"), nil
	})

	if err != nil {
		return nil, err
	}

	if err=claims.Valid(); err != nil {
		return nil,err
	}
	return claims.User,err
}