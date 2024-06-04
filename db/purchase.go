package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/IstiyakRiyad/restaurant-service/internal/restaurant"
)

func (db *DataBase) CreateManyPurchase(purchases []restaurant.Purchase) error {
	var sqlArgs []any
	counter := 0
	sqlCommand := "insert into purchases (amount, purchase_time, restaurant_id, user_id, menu_id) values "

	for i, purchase := range purchases {
		sqlCommand += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", counter + 1, counter + 2, counter + 3, counter + 4, counter + 5)
		sqlArgs = append(sqlArgs, purchase.TransactionAmount, purchase.TransactionDate, purchase.RestaurantId, purchase.UserId, purchase.MenuId)
		
		// If the item is not finised or not reached to 1000
		if (i % 1000) != 0 && i != len(purchases) - 1 {
			sqlCommand += ","
			counter += 5
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
		sqlCommand = "insert into purchases (amount, purchase_time, restaurant_id, user_id, menu_id) values "
	}

	return nil
}


// Purchase Error creates an order for an album and returns the new order ID.
func (db *DataBase) CreatePurchase(ctx context.Context, userId, dishId int) error {
    tx, err := db.Client.BeginTx(ctx, nil)
    if err != nil {
        return fmt.Errorf("Purchase Error: %v", err)
    }

    // Defer a rollback in case anything fails.
    defer tx.Rollback()

	// Query the user & dish data
	var menu restaurant.Menu
	if err := tx.QueryRowContext(ctx, "select * from menus where menus.id = $1;", dishId).Scan(
		&menu.ID, &menu.Name, &menu.Price, &menu.RestaurantId); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Menu not found")
		}
		return fmt.Errorf("Purchase Error: %v",  err)
	}

	var user restaurant.User
	if err := tx.QueryRowContext(ctx, "select * from users where users.id = $1;", userId).Scan(
		&user.ID, &user.Name, &user.CashBalance); err != nil {
		fmt.Println("ehHellowrold")
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("Purchase Error: %v",  err)
	}

	if user.CashBalance < menu.Price {
		return fmt.Errorf("Not enough balance")
	}

    _, err = tx.ExecContext(ctx, "update users set cash_balance = cash_balance - $1 where id = $2;", menu.Price, user.ID)
    if err != nil {
        return fmt.Errorf("Purchase Error: %v", err)
    }

    _, err = tx.ExecContext(ctx, "update restaurants set cash_balance = cash_balance + $1 where id = $2;", menu.Price, menu.RestaurantId)
    if err != nil {
        return fmt.Errorf("Purchase Error: %v", err)
    }

    // Create a new row in the album_order table.
    _, err = tx.ExecContext(ctx, "insert into purchases (amount, restaurant_id, user_id, menu_id) values ($1, $2, $3, $4);",
        menu.Price, menu.RestaurantId, user.ID, menu.ID)
    if err != nil {
        return fmt.Errorf("Purchase Error: %v", err)
    }

    // Commit the transaction.
    if err = tx.Commit(); err != nil {
        return fmt.Errorf("Purchase Error: %v", err)
    }

    // Return the order ID.
    return nil
}
