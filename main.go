package main

/// Go fmt import
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"famtree-api/config"
	"famtree-api/model"

	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func getPort() string {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	return port

}

// Go main function
func main() {
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/families/", GetFamilies).Methods("GET")
	router.HandleFunc("/check-id-family/{family-id}", CheckIsFamilyIdAvailable).Methods("GET")

	// Create a movie
	// router.HandleFunc("/movies/", CreateMovie).Methods("POST")

	// // Delete a specific movie by the movieID
	// router.HandleFunc("/movies/{movieid}", DeleteMovie).Methods("DELETE")

	// // Delete all movies
	// router.HandleFunc("/movies/", DeleteMovies).Methods("DELETE")

	// serve the app
	port := getPort()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("Failed starting http server: ", err)
	}
}

func GetFamilies(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	PrintMessage("Getting books...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT id,name FROM famtree.family")

	// check errors
	CheckErr(err)

	// var response []JsonResponse
	var family []model.Family

	// Foreach movie
	for rows.Next() {
		var id string
		var name string

		err = rows.Scan(&id, &name)

		// check errors
		CheckErr(err)

		family = append(family, model.Family{FamilyID: id, FamilyName: name})
	}

	var response = JsonResponse{Type: "success", Data: family, Message: "Yeay"}

	json.NewEncoder(w).Encode(response)
}

func CheckIsFamilyIdAvailable(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

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

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

type CheckResult struct {
	Type    string `json:"type"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type JsonResponse struct {
	Type    string         `json:"type"`
	Data    []model.Family `json:"data"`
	Message string         `json:"message"`
}
