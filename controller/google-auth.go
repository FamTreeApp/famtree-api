package controller

import (
	"encoding/json"
	"famtree-api/model"
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

	var result = model.User_Account{User_Account_id: user.UserID, User_Account_name: user.Name}
	json.NewEncoder(w).Encode(result)
}
