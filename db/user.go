package db

import (
	"fmt"
	"context"
	"database/sql"

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


func (db *DataBase) GetUserById(ctx context.Context, id int) (*restaurant.User, error) {
	sqlCommand := `select id, name, cash_balance from users 
		where id = $1;` 

	var user restaurant.User;
	if err := db.Client.QueryRowContext(ctx, sqlCommand, id).Scan(&user.ID, &user.Name, &user.CashBalance); err != nil {
		if err == sql.ErrNoRows {
            return &user, nil
        }
        return &user, err
	}

	return &user, nil
}

func (db *DataBase) GetUsers(ctx context.Context) ([]restaurant.User, error) {
	sqlCommand := `select * from users` 
	
	rows, err := db.Client.QueryContext(ctx, sqlCommand) 
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	users := []restaurant.User{}

    for rows.Next() {
        var user restaurant.User;

        if err := rows.Scan(&user.ID, &user.Name, &user.CashBalance); err != nil {
            return users, err
        }

        users = append(users, user)
    }
    if err = rows.Err(); err != nil {
        return users, err
    }

    return users, nil
}

