package http

import (
	"fmt"
	"net/http"
	"time"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)


type RestaurantService interface {
	GetRestaurantsByDate(time.Time) ([]restaurant.Restaurant, error)
	GetRestaurants(restaurant.RestaurantQuery) ([]restaurant.Restaurant, error)
	GetRestaurantById(int) (* restaurant.Restaurant, error)
	SearchRestaurant(string) ([]restaurant.Restaurant, error)
	SearchDish(string) ([]restaurant.Menu, error)
	Purchase(context.Context, int, int) error
	GetUsers() ([]restaurant.User, error)
	GetUserById(int) (restaurant.User, error)
}

type Handler struct {
	Service RestaurantService
	Router *gin.Engine
	Server *http.Server
}


func NewHandler(service RestaurantService) *Handler {
	hander := &Handler{
		Service: service,
	}

	// gin.SetMode(gin.ReleaseMode)
	hander.Router = gin.Default()
	hander.mapRoute()

	serverAddr := fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
	hander.Server = &http.Server{
		Addr: serverAddr,
		Handler: hander.Router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return hander
}


func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}


func (h *Handler) mapRoute() {
	v1 := h.Router.Group("/api/v1")

	// Restaurant Group
	restaurantGroup := v1.Group("/restaurant")
	restaurantGroup.GET("/", h.GetRestaurants)
	restaurantGroup.GET("/datetime", h.GetRestaurantsByDate)
	restaurantGroup.GET("/:id", h.GetRestaurantById)
	
	// User Group
	userGroup := v1.Group("/user")
	userGroup.GET("/", h.GetUsers)
	userGroup.GET("/:id", h.GetUserById)

	// Purchase
	purchaseGroup := v1.Group("/purchase")
	purchaseGroup.POST("/", h.Purchase)

	// Sesarch
	searchGroup := v1.Group("/search")
	searchGroup.GET("/", h.Search)

	// Not Found
	h.Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Route Not Found",
		})
	})
}




