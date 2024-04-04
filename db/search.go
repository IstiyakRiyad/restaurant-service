package db

import (
	"fmt"

	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)


func (db *DataBase) SearchRestaurant(searchQuery string) ([]restaurant.Restaurant, error) {
	sqlCommand := `select id, name, cash_balance,
				ts_rank(to_tsvector('english', name), websearch_to_tsquery($1)) as rank
				from restaurants
				where to_tsvector('english', name) @@ websearch_to_tsquery($2)
				order by rank desc;` 

	rows, err := db.Client.Query(sqlCommand, searchQuery, searchQuery) 
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	restaurants := []restaurant.Restaurant{}

    for rows.Next() {
        var restaurant restaurant.Restaurant;
		var rank float64

        if err := rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.CashBalance, &rank); err != nil {
			fmt.Println(err)
            return restaurants, err
        }

        restaurants = append(restaurants, restaurant)
    }
    if err = rows.Err(); err != nil {
        return restaurants, err
    }

    return restaurants, nil
}

func (db *DataBase) SearchDish(searchQuery string) ([]restaurant.Menu, error) {
	sqlCommand := `select id, name, price, restaurant_id, 
				ts_rank(to_tsvector('english', name), websearch_to_tsquery($1)) as rank
				from menus
				where to_tsvector('english', name) @@ websearch_to_tsquery($2)
				order by rank desc;` 

	rows, err := db.Client.Query(sqlCommand, searchQuery, searchQuery) 
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	menus := []restaurant.Menu{}

    for rows.Next() {
        var menu restaurant.Menu
		var rank float64

        if err := rows.Scan(&menu.ID, &menu.Name, &menu.Price, &menu.RestaurantId, &rank); err != nil {
            return menus, err
        }

        menus = append(menus, menu)
    }
    if err = rows.Err(); err != nil {
        return menus, err
    }

    return menus, nil
}

