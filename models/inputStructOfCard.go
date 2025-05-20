package models

import _ "github.com/go-playground/validator/v10"

type CoffeeCardInput struct {
	Id          string  `db:"id" json:"id"`
	Img         string  `db:"img" json:"img" validate:"required,url"`
	CoffeeName  string  `db:"coffee_name" json:"coffeeName" validate:"required,min=3"`
	Country     string  `db:"country" json:"country" validate:"required,min=3"`
	Price       float64 `db:"price" json:"price" validate:"required,gt=0"`
	Recommended bool    `db:"recommended" json:"recommended"`
	Description string  `db:"description" json:"description" validate:"required,min=3"`
}
