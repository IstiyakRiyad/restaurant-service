package db

import (
	"database/sql"
	"fmt"
	"time"

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

func (db *DataBase) GetRestaurantsByDate(week string, timeUnix time.Time) ([]restaurant.Restaurant, error) {
	sqlCommand := `select restaurants.* from opening_hours join restaurants 
					on restaurants.id = opening_hours.restaurant_id
					where opening_hours.day = $1 and $2 between opening_hours.start_time and opening_hours.end_time` 
	
	
	rows, err := db.Client.Query(sqlCommand, week, timeUnix) 
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	restaurants := []restaurant.Restaurant{}

    for rows.Next() {
        var restaurant restaurant.Restaurant;

        if err := rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.CashBalance); err != nil {
            return restaurants, err
        }

        restaurants = append(restaurants, restaurant)
    }
    if err = rows.Err(); err != nil {
        return restaurants, err
    }

    return restaurants, nil
}

func (db *DataBase) GetRestaurantsMoreThan(limit, baseCount int, baseType string, minPrice, maxPrice float64) ([]restaurant.Restaurant, error) {
	sqlCommand := `select id, name, cash_balance from (select count(*) as no_of_menus, restaurant_id from menus 
					where price between $1 and $2 group by restaurant_id) as cout_table 
					join restaurants on restaurants.id = cout_table.restaurant_id
					where no_of_menus > $3 order by restaurants.name limit $4;` 

	rows, err := db.Client.Query(sqlCommand, minPrice, maxPrice, baseCount, limit) 
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	restaurants := []restaurant.Restaurant{}

    for rows.Next() {
        var restaurant restaurant.Restaurant;

        if err := rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.CashBalance); err != nil {
            return restaurants, err
        }

        restaurants = append(restaurants, restaurant)
    }
    if err = rows.Err(); err != nil {
        return restaurants, err
    }

    return restaurants, nil
}

func (db *DataBase) GetRestaurantsLessThan(limit, baseCount int, baseType string, minPrice, maxPrice float64) ([]restaurant.Restaurant, error) {
	sqlCommand := `select id, name, cash_balance from (select count(*) as no_of_menus, restaurant_id from menus 
			where price between $1 and $2 group by restaurant_id) as cout_table 
			join restaurants on restaurants.id = cout_table.restaurant_id
			where no_of_menus < $3 order by restaurants.name limit $4;` 

	rows, err := db.Client.Query(sqlCommand, minPrice, maxPrice, baseCount, limit) 
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	restaurants := []restaurant.Restaurant{}

    for rows.Next() {
        var restaurant restaurant.Restaurant;

        if err := rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.CashBalance); err != nil {
            return restaurants, err
        }

        restaurants = append(restaurants, restaurant)
    }
    if err = rows.Err(); err != nil {
        return restaurants, err
    }

    return restaurants, nil
}

func (db *DataBase) GetRestaurantById(id int) (*restaurant.Restaurant, error) {
	sqlCommand := `select id, name, cash_balance from restaurants 
		where restaurants.id = $1;` 

	var restaurant restaurant.Restaurant;
	if err := db.Client.QueryRow(sqlCommand, id).Scan(&restaurant.ID, &restaurant.Name, &restaurant.CashBalance); err != nil {
		if err == sql.ErrNoRows {
            return &restaurant, nil
        }
        return &restaurant, err
	}

	return &restaurant, nil
}

