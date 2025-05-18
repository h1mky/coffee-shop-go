package models

type CoffeeCard struct {
	Id          string  `json:"id"`
	Img         string  `json:"img"`
	CoffeeName  string  `json:"coffeeName"`
	Country     string  `json:"country"`
	Price       float64 `json:"price"`
	Recommended bool    `json:"recommended"`
	Description string  `json:"description"`
}
