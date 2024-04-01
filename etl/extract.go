package etl

import (
	"encoding/json"
	"log"
	"os"
)

var restaurant_with_menu = "restaurantData/restaurant_with_menu.json"
var users_with_purchase_history = "restaurantData/users_with_purchase_history.json"

type Restaurant struct {
	Name			string		`json:"restaurantName"`
	OpeningHours	string		`json:"openingHours"`
	CashBalance		float64		`json:"cashBalance"`
	Menu			[]struct{
		DishName		string		`json:"dishName"`
		Price			float64		`json:"price"`
	}							`json:"menu"`
}

type User struct {
	ID					int			`json:"id"`
	Name				string		`json:"name"`
	CashBalance			float64		`json:"cashBalance"`
	PurchaseHistory		[]struct{
		DishName			string		`json:"dishName"`
		RestaurantName		string		`json:"restaurantName"`
		TransactionAmount	float64		`json:"transactionAmount"`
		TransactionDate		string		`json:"transactionDate"`
	}								`json:"purchaseHistory"`
}

func extractRestaurantData() ([]Restaurant) {
	file, err := os.Open(restaurant_with_menu)
	if err != nil {
		log.Fatal("Faild to open restaurant file: ", err)
	}

	dec := json.NewDecoder(file)

	var restaurents []Restaurant

	if err := dec.Decode(&restaurents); err != nil {
		log.Fatal("Faild to decode restaurant data: ", err)
	}

	return restaurents
}

func extractUserData() ([]User) {
	file, err := os.Open(users_with_purchase_history)
	if err != nil {
		log.Fatal("Faild to open users file: ", err)
	}

	dec := json.NewDecoder(file)

	var users []User

	if err := dec.Decode(&users); err != nil {
		log.Fatal("Faild to decode users data: ", err)
	}

	return users
}

func ExtractData() ([]Restaurant, []User) {
	restaurants := extractRestaurantData()
	users := extractUserData()

	return restaurants, users
}

