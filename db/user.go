package db

import (
	"fmt"

	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)

func (db *DataBase) CreateManyUser(users []restaurant.User) error {
	var sqlArgs []any
	counter := 0
	sqlCommand := "insert into users (id, name, cash_balance) values "

	for i, user := range users {
		sqlCommand += fmt.Sprintf("($%d, $%d, $%d)", counter + 1, counter + 2, counter + 3)
		sqlArgs = append(sqlArgs, user.ID, user.Name, user.CashBalance)
		
		// If the item is not finised or not reached to 1000
		if (i % 1000) != 0 && i != len(users) - 1 {
			sqlCommand += ","
			counter += 3
			continue
		}
		
		// Create all the items
		_, err := db.Client.Exec(sqlCommand + ";", sqlArgs...) 
		if err != nil {
			return fmt.Errorf("Error Creating user: %v\n", err)
		}

		// Set for restart the formating
		sqlArgs = []any{}
		counter = 0
		sqlCommand = "insert into users (id, name, cash_balance) values "
	}

	return nil
}
