package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"os"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.Handle("/postform",jwtMiddleware.Handler(Loginform))
	r.Handle("/get-token",GetTokenHandler).Methods("GET")
	fmt.Println("Server is listening...")
	//print every request on screen
	http.ListenAndServe(":8181", handlers.LoggingHandler(os.Stdout, r))
}
//Loginform sends postform for get data
var Loginform = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	pass := r.FormValue("userpassword")
	fmt.Fprintf(w, "Enter name: %s enter password: %s", name, pass)

})
var mySigningKey = []byte("secret")

//GetTokenHandler install get set and sends to client token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//create new token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//Install set argumetns for token
	claims["admin"] = true
	claims["name"] = "Nikolai"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims
	//sign token our secret key
	tokenString, _ := token.SignedString(mySigningKey)
	//give token to client
	w.Write([]byte(tokenString))
})
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
