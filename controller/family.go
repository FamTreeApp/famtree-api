package controller

import (
	"encoding/json"
	"famtree-api/db"
	"famtree-api/model"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFamilies(w http.ResponseWriter, r *http.Request) {
	db := db.SetupDB()

	rows, err := db.Query("SELECT id,name FROM famtree.family")
	CheckErr(err)

	var family []model.Family

	for rows.Next() {
		var id string
		var name string

		err = rows.Scan(&id, &name)
		CheckErr(err)

		family = append(family, model.Family{FamilyID: id, FamilyName: name})
	}

	var response = JsonResponse{Type: "success", Data: family, Message: "Yeay"}

	json.NewEncoder(w).Encode(response)
}

func CheckIsFamilyIdAvailable(w http.ResponseWriter, r *http.Request) {
	db := db.SetupDB()

	params := mux.Vars(r)
	family_id := params["family-id"]

	rows, err := db.Query(fmt.Sprintf("SELECT COUNT(*) FROM famtree.FAMILY WHERE id = '%v';", family_id))
	CheckErr(err)

	var result = "false"
	for rows.Next() {
		var count int
		err = rows.Scan((&count))
		CheckErr(err)
		if count == 0 {
			result = "true"
		}
	}

	var response = CheckResult{"Success", result, ""}

	json.NewEncoder(w).Encode(response)
}

type JsonResponse struct {
	Type    string         `json:"type"`
	Data    []model.Family `json:"data"`
	Message string         `json:"message"`
}

type CheckResult struct {
	Type    string `json:"type"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
