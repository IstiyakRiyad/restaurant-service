package etl

import (
	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)

type StoreETL interface {
	CreateManyUser([]restaurant.User) error
	CreateManyRestaurant([]restaurant.Restaurant) error
	CreateManyMenu([]restaurant.Menu) error
	CreateManyOpeningHour([]restaurant.OpeningHour) error
	CreateManyPurchase([]restaurant.Purchase) error
}

type RestaurantETL struct {
	Stote			StoreETL
	restaurantFile	string
	userFile		string
}

func NewETL(store StoreETL) *RestaurantETL {
	return &RestaurantETL{
		Stote: store,
		restaurantFile: "restaurantData/restaurant_with_menu.json",
		userFile: "restaurantData/users_with_purchase_history.json",
	}
}


func (etl *RestaurantETL) Start() {
	// Extract the data from json files
	restaurants, users := etl.extractData()

	// Transform the data
	formatedUsers, formatedRestaurants, formatedMenus, formatedOpeningHours, formatedPurchases := etl.transformData(restaurants, users)

	// Load the data to the database
	etl.loadData(formatedUsers, formatedRestaurants, formatedMenus, formatedOpeningHours, formatedPurchases)
}



