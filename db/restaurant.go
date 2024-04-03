package db

import (
	"fmt"

	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)

func (db *DataBase) CreateManyRestaurant(restaurants []restaurant.Restaurant) error {
	var sqlArgs []any
	counter := 0
	sqlCommand := "insert into restaurants (id, name, cash_balance) values "

	for i, restaurant := range restaurants {
		sqlCommand += fmt.Sprintf("($%d, $%d, $%d)", counter + 1, counter + 2, counter + 3)
		sqlArgs = append(sqlArgs, restaurant.ID, restaurant.Name, restaurant.CashBalance)
		
		// If the item is not finised or not reached to 1000
		if (i % 1000) != 0 && i != len(restaurants) - 1 {
			sqlCommand += ","
			counter += 3
			continue
		}
		
		// Create all the items
		_, err := db.Client.Exec(sqlCommand + ";", sqlArgs...) 
		if err != nil {
			return fmt.Errorf("Error Creating Restaurant: %v\n", err)
		}

		// Set for restart the formating
		sqlArgs = []any{}
		counter = 0
		sqlCommand = "insert into restaurants (id, name, cash_balance) values "
	}

	return nil
}

func (db *DataBase) CreateManyMenu(menus []restaurant.Menu) error {
	var sqlArgs []any
	counter := 0
	sqlCommand := "insert into menus (id, name, price, restaurant_id) values "

	for i, menu := range menus {
		sqlCommand += fmt.Sprintf("($%d, $%d, $%d, $%d)", counter + 1, counter + 2, counter + 3, counter + 4)
		sqlArgs = append(sqlArgs, menu.ID, menu.Name, menu.Price, menu.RestaurantId)
		
		// If the item is not finised or not reached to 1000
		if (i % 1000) != 0 && i != len(menus) - 1 {
			sqlCommand += ","
			counter += 4
			continue
		}
		
		// Create all the items
		_, err := db.Client.Exec(sqlCommand + ";", sqlArgs...) 
		if err != nil {
			return fmt.Errorf("Error Creating Menu: %v\n", err)
		}

		// Set for restart the formating
		sqlArgs = []any{}
		counter = 0
		sqlCommand = "insert into menus (id, name, price, restaurant_id) values "
	}

	return nil
}

func (db *DataBase) CreateManyOpeningHour(openingHours []restaurant.OpeningHour) error {
	var sqlArgs []any
	counter := 0
	sqlCommand := "insert into opening_hours (id, day, start_time, end_time, restaurant_id) values "

	for i, openingHour := range openingHours {
		sqlCommand += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", counter + 1, counter + 2, counter + 3, counter + 4, counter + 5)
		sqlArgs = append(sqlArgs, openingHour.ID, openingHour.Day, openingHour.StartTime, openingHour.EndTime, openingHour.RestaurantId)
		
		// If the item is not finised or not reached to 1000
		if (i % 1000) != 0 && i != len(openingHours) - 1 {
			sqlCommand += ","
			counter += 5
			continue
		}
		
		// Create all the items
		_, err := db.Client.Exec(sqlCommand + ";", sqlArgs...) 
		if err != nil {
			return fmt.Errorf("Error Creating Opening Hour: %v\n", err)
		}

		// Set for restart the formating
		sqlArgs = []any{}
		counter = 0
		sqlCommand = "insert into opening_hours (id, day, start_time, end_time, restaurant_id) values "
	}

	return nil
}

