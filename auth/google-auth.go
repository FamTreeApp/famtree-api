package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func Login(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func LoginCallback(w http.ResponseWriter, r *http.Request) {

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, r)
		return
	}
	json.NewEncoder(w).Encode(user)
}
