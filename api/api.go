package api

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/arnaudbreton/whatisyourstack/models"
	"encoding/json"
	"log"
	// "fmt"
)

func NewApp(db *gorm.DB) *martini.ClassicMartini {
	m := martini.Classic()

	m.Get("/stacks", GetStacks)
	m.Post("/stacks", PostStack)
	m.Map(db)
  	return m
}

func PostStack(res http.ResponseWriter, req *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(req.Body)
   	var stackApi models.StackApi   
   	err := decoder.Decode(&stackApi)
	if err != nil {
		res.WriteHeader(400)
    }

	stack, err := models.NewStackFromApi(&stackApi)
	log.Println(stack)

	if err != nil {
		res.WriteHeader(400)
	}
	db.Save(stack)
    header := res.Header()
	header.Add("Content-Type", "application/json")
	// stack, err := json.Unmarshal(req.Body)

	// if err == nil {
	// 	header := res.Header()
	// 	header.Add("Content-Type", "application/json")
	// 	// res.Write(responseJSON)
	// } else {
	// 	res.WriteHeader(400)
	// }
}

func GetStacks(res http.ResponseWriter, req *http.Request, db *gorm.DB) {
	stacks := make([]models.Stack, 0)
	db.Find(&stacks)

	response := make(map[string][]models.Stack)
	response["stacks"] = stacks
	responseJSON, err := json.Marshal(response)

	if err == nil {
		header := res.Header()
		header.Add("Content-Type", "application/json")
		res.Write(responseJSON)
	} else {
		res.WriteHeader(500)
	}
}