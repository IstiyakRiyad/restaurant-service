package db

import (
	"fmt"

	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)

func (db *DataBase) CreateManyPurchase(purchases []restaurant.Purchase) error {
	var sqlArgs []any
	counter := 0
	sqlCommand := "insert into purchases (id, amount, purchase_time, restaurant_id, user_id, menu_id) values "

	for i, purchase := range purchases {
		sqlCommand += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", counter + 1, counter + 2, counter + 3, counter + 4, counter + 5, counter + 6)
		sqlArgs = append(sqlArgs, purchase.ID, purchase.TransactionAmount, purchase.TransactionDate, purchase.RestaurantId, purchase.UserId, purchase.MenuId)
		
		// If the item is not finised or not reached to 1000
		if (i % 1000) != 0 && i != len(purchases) - 1 {
			sqlCommand += ","
			counter += 6
			continue
		}
		
		// Create all the items
		_, err := db.Client.Exec(sqlCommand + ";", sqlArgs...) 
		if err != nil {
			return fmt.Errorf("Error Creating Purchage History: %v\n", err)
		}

		// Set for restart the formating
		sqlArgs = []any{}
		counter = 0
		sqlCommand = "insert into purchases (id, amount, purchase_time, restaurant_id, user_id, menu_id) values "
	}

	return nil
}

