package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type session struct {
	userid string
	cookie string
}

var users map[string]string
var sessions map[int]session
var sessionCount int

func main() {
	http.HandleFunc("/", serveFiles)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	fmt.Printf("Serving on port %d", 8080)
	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("../build")).ServeHTTP(w, r)
}

func register(w http.ResponseWriter, r *http.Request) {
	// is email properly formated? NO, send error msg (404?)
	// Is password valid? NO, send error msg (404?)
	// userid in database? NO, send error msg (404?)
	// hash password
	// store userid and pw in users
	// send okay msg 200
}

func login(w http.ResponseWriter, r *http.Request) {
	// cookie set?
	//    check if valid?
	//       No, errase cookie, then continue with flow
	//       valid, send already logged in msg
	// Does user id exist?
	//		No, send "email or password is incorrect"
	// Get password hash for userid
	// Does it match the password received
	// generate a random secure cookie
	// Generate unique session
	// store cookie and userid in session
	// send cookie and session to client
}
