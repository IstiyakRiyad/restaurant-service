package etl

import (
	"encoding/json"
	"log"
	"os"
)

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

func (etl *RestaurantETL) extractRestaurantData() ([]Restaurant) {
	file, err := os.Open(etl.restaurantFile)
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

func (etl *RestaurantETL) extractUserData() ([]User) {
	file, err := os.Open(etl.userFile)
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

func (etl *RestaurantETL) extractData() ([]Restaurant, []User) {
	restaurants := etl.extractRestaurantData()
	users := etl.extractUserData()

	return restaurants, users
}

