package etl

import (
	"time"

	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/utils"
	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)


func (etl *RestaurantETL) transformData(restaurants []Restaurant, users []User) ( []restaurant.User, []restaurant.Restaurant, []restaurant.Menu, []restaurant.OpeningHour, []restaurant.Purchase ){
	var (
		formatedUsers []restaurant.User
		formatedRestaurants []restaurant.Restaurant
		formatedMenus []restaurant.Menu
		formatedOpeningHours []restaurant.OpeningHour
		formatedPurchases []restaurant.Purchase
	)
	restaurantCount := 1
	menuCount := 1
	openingHourCount := 1
	purchaseCount := 1

	for _, restaurantVal := range restaurants {
		formatedRestaurants = append(formatedRestaurants, restaurant.Restaurant{
			ID: restaurantCount,
			Name: restaurantVal.Name,
			CashBalance: restaurantVal.CashBalance,
		})

		for _, menu := range restaurantVal.Menu {
			formatedMenus = append(formatedMenus, restaurant.Menu{
				ID: menuCount,
				Name: menu.DishName,
				Price: menu.Price,
				RestaurantId: restaurantCount,
			})

			menuCount++
		}

		for _, schedule := range utils.WeeklyScheduleDecoder(restaurantVal.OpeningHours){
			formatedOpeningHours = append(formatedOpeningHours, restaurant.OpeningHour{
				ID: openingHourCount,
				Day: schedule.Day,
				StartTime: schedule.StartTime,
				EndTime: schedule.EndTime,
				RestaurantId: restaurantCount,
			})

			openingHourCount++
		}

		restaurantCount++
	}

	for _, user := range users {
		formatedUsers = append(formatedUsers, restaurant.User{
			ID: user.ID,
			Name: user.Name,
			CashBalance: user.CashBalance,
		})
		
		// 02/01/2006 03:04 PM
		for _, purchase := range user.PurchaseHistory {
			parsedTime, err := time.Parse("02/01/2006 03:04 PM",  purchase.TransactionDate)
			if err != nil {
				parsedTime = time.Now()
			}

			// Find restaurantId & the menuId
			var (
				restaurantId	int
				menuId			int
			)
			for _, restaurant := range formatedRestaurants {
				if restaurant.Name != purchase.RestaurantName { continue }

				for _, menu := range formatedMenus {
					if menu.Name != purchase.DishName { continue }
					restaurantId = restaurant.ID
					menuId = menu.ID
				}
			}

			formatedPurchases = append(formatedPurchases, restaurant.Purchase{
				ID: purchaseCount,
				TransactionAmount: purchase.TransactionAmount,
				TransactionDate: parsedTime,
				UserId: user.ID,
				RestaurantId: restaurantId,
				MenuId: menuId,
			})

			purchaseCount++
		}
	}

	return formatedUsers, formatedRestaurants, formatedMenus, formatedOpeningHours, formatedPurchases 
}






