package restaurant

import (
	"time"
)

type RestaurantStore interface {
}

type RestaurantService struct {
	Store RestaurantStore
}

type RestaurantQuery struct {
	Limit		int		`form:"limit,default=10"`
	BaseCount   int		`form:"base_count,default=1"`
	BaseType	string	`form:"base_type,default=more"`
	MinPrice	float64	`form:"min_price"`
	MaxPrice	float64	`form:"max_price"`
}


func NewRestaurantService(store RestaurantStore) *RestaurantService {

	return &RestaurantService{
		Store: store,
	}
}

func (rs *RestaurantService) GetRestaurantsByDate(dateTime time.Time) ([]Restaurant, error){
	return []Restaurant{}, nil
}

func (rs *RestaurantService) GetRestaurants(query RestaurantQuery) ([]Restaurant, error){
	return []Restaurant{}, nil
}

func (rs *RestaurantService) GetRestaurantById(id int) ([]Restaurant, error){
	return []Restaurant{}, nil
}

func (rs *RestaurantService) SearchRestaurant(searchQuery string) ([]Restaurant, error){
	return []Restaurant{}, nil
}

func (rs *RestaurantService) SearchDish(searchQuery string) ([]Menu, error){
	return []Menu{}, nil

}

func (rs *RestaurantService) Purchase(userId int, dishId int) error{
	return nil
}

func (rs *RestaurantService) GetUsers() ([]User, error){
	return []User{}, nil
}

func (rs *RestaurantService) GetUserById(id int) (User, error){
	return User{}, nil
}




