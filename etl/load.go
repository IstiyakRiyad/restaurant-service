package etl

import (
	"log"

	"github.com/IstiyakRiyad/restaurant-service/internal/restaurant"
)


func (etl *RestaurantETL) loadData(
	formatedUsers []restaurant.User,
	formatedRestaurants []restaurant.Restaurant,
	formatedMenus []restaurant.Menu,
	formatedOpeningHours []restaurant.OpeningHour,
	formatedPurchases []restaurant.Purchase,
) {
	// Create the restaurans 
	if err := etl.Stote.CreateManyRestaurant(formatedRestaurants); err != nil {
		log.Fatal(err);
	}

	// Create the menu
	if err := etl.Stote.CreateManyMenu(formatedMenus); err != nil {
		log.Fatal(err);
	}

	// Create the opening hours
	if err := etl.Stote.CreateManyOpeningHour(formatedOpeningHours); err != nil {
		log.Fatal(err);
	}

	// Create the users
	if err := etl.Stote.CreateManyUser(formatedUsers); err != nil {
		log.Fatal(err);
	}

	// Create the formatedPurchases
	if err := etl.Stote.CreateManyPurchase(formatedPurchases); err != nil {
		log.Fatal(err);
	}
}


