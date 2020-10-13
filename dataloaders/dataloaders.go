package dataloaders

import (
	. "../models"
	. "../utils"
	"encoding/json"
)

func LoadUsers() []User {
	bytes, _ := ReadFile("../json/data.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterests() []Interest {
	bytes, _ := ReadFile("../json/data.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterestMapping() []InterestMapping {
	bytes, _ := ReadFile("../json/userInterestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
