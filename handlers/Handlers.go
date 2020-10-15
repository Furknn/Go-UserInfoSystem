package handlers

import (
	. "../dataloaders"
	. "../models"
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	//Mock Data TODO:Replace
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}

	users := LoadUsers()
	interests := LoadInterests()
	interestMappings := LoadInterestMapping()
	var newUsers []User

	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interest.ID == interestMapping.InterestID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
