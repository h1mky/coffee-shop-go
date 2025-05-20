package db

import "coffee-shop/models"

func ToParams(id int, card models.CoffeeCard) map[string]interface{} {

	return map[string]interface{}{
		"id":          id,
		"img":         card.Img,
		"coffee_name": card.CoffeeName,
		"country":     card.Country,
		"price":       card.Price,
		"recommended": card.Recommended,
		"description": card.Description,
	}
}
