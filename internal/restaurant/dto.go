package restaurant

import (
	"time"
)

type User struct {
	ID			int
	Name		string
	CashBalance	float64
}

type Restaurant struct {
	ID				int
	Name			string
	CashBalance		float64
}

type Menu struct {
	ID				int
	Name			string
	Price			float64
	RestaurantId	int
}

type OpeningHour struct {
	ID				int
	Day				string
	StartTime		time.Time
	EndTime			time.Time
	RestaurantId	int
}


type Purchase struct {
	ID					int
	TransactionAmount	float64
	TransactionDate		time.Time
	RestaurantId		int
	UserId				int
	MenuId				int
}

