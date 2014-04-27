package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "github.com/arnaudbreton/whatisyourstack/models"
)

func NewDB() *gorm.DB {
    db, err := gorm.Open("postgres", "user=whatisyourstack dbname=whatisyourstack sslmode=disable")

    db.SingularTable(true)

    if err != nil {
        panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
    }

    return &db
}

func Migrate(db *gorm.DB) {
	db.CreateTable(models.Language{})
	db.CreateTable(models.Technology{})
	db.CreateTable(models.Company{})
	db.CreateTable(models.Stack{})
}