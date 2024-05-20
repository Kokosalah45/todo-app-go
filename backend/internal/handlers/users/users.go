package users

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var FakeUsers = []User{
	{ID: "1", FirstName: "John", LastName: "Doe"},
	{ID: "2", FirstName: "Jane", LastName: "Doe"},
	{ID: "3", FirstName: "John", LastName: "Smith"},
	{ID: "4", FirstName: "Jane", LastName: "Smith"},
	{ID: "5", FirstName: "John", LastName: "Johnson"},
}


func ListUsers(w http.ResponseWriter , r *http.Request){

	users := FakeUsers

	encoder := json.NewEncoder(w)


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := encoder.Encode(&users)

	if err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}

}
func GetUserByID(w http.ResponseWriter , r *http.Request){
	id := r.PathValue("id")
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, user := range FakeUsers {
		if user.ID == id {
			err := encoder.Encode(&user)
			if err != nil {
				http.Error(w, "Failed to encode user", http.StatusInternalServerError)
				return
			}
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
