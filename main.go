package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/get-token", GetTokenHandler).Methods("GET")
	r.Handle("/login", jwtMiddleware.Handler(LoginHandler))
	//print every request on screen
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
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

const loginForm string = `<html>
	<head></head><body>
	<form action="login" method="post">
		<label for="uname"><b>Username</b></label>
		<input type="text" placeholder="Enter Username" name="uname" required>

		<label for="psw"><b>Password</b></label>
		<input type="password" placeholder="Enter Password" name="password" required>

		<button type="submit">Login</button>
	</form>
	</body>
	</html>`

const loginFormRu string = `<html>
	<head></head><body>
	<form action="login" method="post">
		<label for="uname"><b>Имя пользователя</b></label>
		<input type="text" placeholder="Enter Username" name="uname" required>

		<label for="psw"><b>Password</b></label>
		<input type="password" placeholder="Enter Password" name="password" required>

		<button type="submit">Login</button>
	</form>
	</body>
	</html>`

const loginFormResult string = `<html>
	<head></head><body>
	Name: %s<p>
	Password: %s
	</body>
	</html>`

// LoginHandler sends login form for GET method
var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, loginForm)
	case "POST":
		r.ParseForm()
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		var uname = r.PostForm.Get("uname")
		var password = r.PostForm.Get("password")
		var responseContent = fmt.Sprintf(loginFormResult, uname, password)
		io.WriteString(w, responseContent)
	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "<html>Bad request</html>")
	}
})