package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)
var MySigningKey = []byte("SECRET_KEY")

func GetJWt() (string, error)  {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorize"] = true
	claims["client"] = "Croissant"
	claims["aud"] = "billing.jingo.io"
	claims["iss"] = "jingo.io"
	claims["exp"] = time.Now().Add(time.Minute *1).Unix()
	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		fmt.Errorf("Something  went wrong: %s", err.Error())
		return "",err
	}
	return tokenString, nil

}
func Index(w http.ResponseWriter, r *http.Request){
	validToken, err := GetJWt()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
}
func handleRequest()  {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
func main(){
	fmt.Println("jwt_service ")
	handleRequest()
	}

