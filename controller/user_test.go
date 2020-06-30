package controller

import (
	"cwm.wiki/web/models"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestJwtGenerateToken(t *testing.T) {

	user := models.User{
		Username: "admin",
		Password: "admin",
	}

	token,err := JwtGenerateToken(&user,time.Duration(time.Minute*20))

	if err != nil {
		log.Fatal("err: " ,err)
	}

	log.Println("token: ",token)
}

func TestJwtParseUser(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTM1MDM3ODYsImlhdCI6MTU5MzUwMjU4NiwidXNlcm5hbWUiOiJhZG1pbiIsInBhc3N3b3JkIjoiYWRtaW4ifQ.wxG7DL7mX8hzJNsbQ42ULlSqmqVyA-JJBLyA-8EnuUI"
	user, err := JwtParseUser(token)

	if err != nil {
		log.Fatal(err)
	}

	if user != nil {
		fmt.Println(*user)
	}

}
