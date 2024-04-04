package restaurant

import (
	"context"
	"time"
)

type RestaurantStore interface {
	GetRestaurantsByDate(context.Context, string, time.Time) ([]Restaurant, error)
	GetRestaurantsLessThan(context.Context, int, int, string, float64, float64) ([]Restaurant, error)
	GetRestaurantsMoreThan(context.Context, int, int, string, float64, float64) ([]Restaurant, error)
	GetRestaurantById(context.Context, int) (*Restaurant, error)
	SearchRestaurant(context.Context, string) ([]Restaurant, error)
	SearchDish(context.Context, string) ([]Menu, error)
	CreatePurchase(context.Context, int, int) error
	GetUsers(context.Context) ([]User, error)
	GetUserById(context.Context, int) (*User, error)
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

func (rs *RestaurantService) GetRestaurantsByDate(ctx context.Context, dateTime time.Time) ([]Restaurant, error){

	week := dateTime.Weekday().String()
	timeUnix := time.Date(1970, 1, 1, dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())

	restaurants, err := rs.Store.GetRestaurantsByDate(ctx, week, timeUnix)
	if err != nil {
		return restaurants, err
	}

	return restaurants, nil
}

func (rs *RestaurantService) GetRestaurants(ctx context.Context, query RestaurantQuery) ([]Restaurant, error){
	if query.BaseType == "more" {
		restaurants, err := rs.Store.GetRestaurantsMoreThan(ctx, query.Limit, query.BaseCount, query.BaseType, query.MinPrice, query.MaxPrice)
		if err != nil {
			return restaurants, err
		}

		return restaurants, nil
	} 

	restaurants, err := rs.Store.GetRestaurantsLessThan(ctx, query.Limit, query.BaseCount, query.BaseType, query.MinPrice, query.MaxPrice)
	if err != nil {
		return restaurants, err
	}

	return restaurants, nil
}

func (rs *RestaurantService) GetRestaurantById(ctx context.Context, id int) (*Restaurant, error){
	restaurant, err := rs.Store.GetRestaurantById(ctx, id)
	if err != nil {
		return restaurant, err
	}

	return restaurant, nil
}

func (rs *RestaurantService) SearchRestaurant(ctx context.Context, searchQuery string) ([]Restaurant, error){
	restaurants, err := rs.Store.SearchRestaurant(ctx, searchQuery)
	if err != nil {
		return restaurants, err
	}

	return restaurants, nil
}

func (rs *RestaurantService) SearchDish(ctx context.Context, searchQuery string) ([]Menu, error){
	dishes, err := rs.Store.SearchDish(ctx, searchQuery)
	if err != nil {
		return dishes, err
	}

	return dishes, nil
}

func (rs *RestaurantService) Purchase(ctx context.Context, userId int, dishId int) error{
	if err := rs.Store.CreatePurchase(ctx, userId, dishId); err != nil {
		return err
	}

	return nil
}

func (rs *RestaurantService) GetUsers(ctx context.Context) ([]User, error){
	users, err := rs.Store.GetUsers(ctx)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (rs *RestaurantService) GetUserById(ctx context.Context, id int) (*User, error){
	user, err := rs.Store.GetUserById(ctx, id)
	if err != nil {
		return user, err
	}

	return user, nil
}




