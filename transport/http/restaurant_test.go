package http_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/IstiyakRiyad/restaurant-service/db"
	"github.com/IstiyakRiyad/restaurant-service/internal/restaurant"
	transportHttp "github.com/IstiyakRiyad/restaurant-service/transport/http"
)


func init() {
	viper.SetConfigType("env")
	viper.SetConfigName("dev")
	viper.AddConfigPath("../..")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}


func TestGetRestaurants(t *testing.T) {
	db, _:= db.NewDatabase()

	service := restaurant.NewRestaurantService(db)
	trasport := transportHttp.NewHandler(service)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/restaurant/", nil)

	trasport.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}


func TestGetRestaurantByDateTime(t *testing.T) {
	db, _ := db.NewDatabase()
	service := restaurant.NewRestaurantService(db)
	trasport := transportHttp.NewHandler(service)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/restaurant/datetime", nil)

	trasport.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}


func TestGetRestaurantById(t *testing.T) {
	db, _ := db.NewDatabase()
	service := restaurant.NewRestaurantService(db)
	trasport := transportHttp.NewHandler(service)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/restaurant/12", nil)

	trasport.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSearch(t *testing.T) {
	db, _ := db.NewDatabase()
	service := restaurant.NewRestaurantService(db)
	trasport := transportHttp.NewHandler(service)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/search/?search_query=beef", nil)

	trasport.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetUsers(t *testing.T) {
	db, _ := db.NewDatabase()
	service := restaurant.NewRestaurantService(db)
	trasport := transportHttp.NewHandler(service)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/", nil)

	trasport.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetUserById(t *testing.T) {
	db, _ := db.NewDatabase()
	service := restaurant.NewRestaurantService(db)
	trasport := transportHttp.NewHandler(service)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/3", nil)

	trasport.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}



