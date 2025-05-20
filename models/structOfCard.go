package models

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type CoffeeCard struct {
	Id          string    `db:"id" json:"id"`
	Img         string    `db:"img" json:"img"`
	CoffeeName  string    `db:"coffee_name" json:"coffeeName"`
	Country     string    `db:"country" json:"country"`
	Price       float64   `db:"price" json:"price"`
	Recommended bool      `db:"recommended" json:"recommended"`
	Description string    `db:"description" json:"description" `
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}
