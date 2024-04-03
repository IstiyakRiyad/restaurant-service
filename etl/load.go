package etl

import (
	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)


func (etl *RestaurantETL) loadData(
	formatedUsers []restaurant.User,
	formatedRestaurants []restaurant.Restaurant,
	formatedMenus []restaurant.Menu,
	formatedOpeningHours []restaurant.OpeningHour,
	formatedPurchases []restaurant.Purchase,
) {
	// Create the restaurans 
	etl.Stote.CreateManyRestaurant(formatedRestaurants)

	// Create the menu
	etl.Stote.CreateManyMenu(formatedMenus)

	// Create the opening hours
	etl.Stote.CreateManyOpeningHour(formatedOpeningHours)

	// Create the users
	etl.Stote.CreateManyUser(formatedUsers)

	// Create the formatedPurchases
	etl.Stote.CreateManyPurchase(formatedPurchases)
}


