package authhandler

import (
	"encoding/json"
	"fmt"
	"github.com/adibur6/bookstoreapi/datahandler"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwa"
	"net/http"
	"time"
)

var Secret = []byte("Ashfaq1234") // Replace with a securely generated key
var TokenAuth *jwtauth.JWTAuth

// InitToken Initialize JWT token authentication
func InitToken() {
	TokenAuth = jwtauth.New(string(jwa.HS256), Secret, nil)
}

// Function signatures
func Login(w http.ResponseWriter, r *http.Request) {
	var cred datahandler.Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	fmt.Println(cred)
	if err != nil {
		http.Error(w, "Can not Decode the data", http.StatusBadRequest)
		return
	}

	password, okay := datahandler.CredList[cred.Username]

	if okay == false {
		http.Error(w, "Username do not exist", http.StatusBadRequest)
		return
	}

	if password != cred.Password {
		http.Error(w, "Password not matching", http.StatusBadRequest)
		return
	}
	et := time.Now().Add(15 * time.Minute)
	_, tokenString, err := TokenAuth.Encode(map[string]interface{}{
		"aud": "ashfaq",
		"exp": et.Unix(),
		// Here few more registered field and also self-driven field can be added
	})
	fmt.Println(tokenString)
	if err != nil {
		http.Error(w, "Can not generate jwt", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: et,
	})
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
}
