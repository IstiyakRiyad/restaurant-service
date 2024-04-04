package restaurant

import (
	"time"
)

type RestaurantStore interface {
	GetRestaurantsByDate(string, time.Time) ([]Restaurant, error)
	GetRestaurantsLessThan(int, int, string, float64, float64) ([]Restaurant, error)
	GetRestaurantsMoreThan(int, int, string, float64, float64) ([]Restaurant, error)
	GetRestaurantById(int) (*Restaurant, error)
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

	week := dateTime.Weekday().String()
	timeUnix := time.Date(1970, 1, 1, dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())

	restaurants, err := rs.Store.GetRestaurantsByDate(week, timeUnix)
	if err != nil {
		return restaurants, nil
	}

	return restaurants, nil
}

func (rs *RestaurantService) GetRestaurants(query RestaurantQuery) ([]Restaurant, error){
	if query.BaseType == "more" {
		restaurants, err := rs.Store.GetRestaurantsMoreThan(query.Limit, query.BaseCount, query.BaseType, query.MinPrice, query.MaxPrice)
		if err != nil {
			return restaurants, nil
		}

		return restaurants, nil
	} 

	restaurants, err := rs.Store.GetRestaurantsLessThan(query.Limit, query.BaseCount, query.BaseType, query.MinPrice, query.MaxPrice)
	if err != nil {
		return restaurants, nil
	}

	return restaurants, nil
}

func (rs *RestaurantService) GetRestaurantById(id int) (*Restaurant, error){
	restaurant, err := rs.Store.GetRestaurantById(id)
	if err != nil {
		return restaurant, nil
	}

	return restaurant, nil
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




