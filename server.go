package main 

import (
	"github.com/arnaudbreton/whatisyourstack/api"
	"github.com/arnaudbreton/whatisyourstack/db"
)

func main() {
  dbInstance := db.NewDB()
  db.Migrate(dbInstance)
  m := api.NewApp(dbInstance)
  m.Run()
}