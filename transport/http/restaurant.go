package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/internal/restaurant"
)

type RestaurantDateQuery struct {
	DateTime	 time.Time	`form:"date_time"`
}

type SearchQuery struct {
	SearchQuery	string	`form:"search_query,default="`
}

func (h *Handler) GetRestaurants(ctx *gin.Context) {
	var query restaurant.RestaurantQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
        ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
        return
    }

	restaurants, err := h.Service.GetRestaurants(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of restaurants",
		"data": restaurants,
	})
}

func (h *Handler) GetRestaurantsByDate(ctx *gin.Context) {
	var query = RestaurantDateQuery{DateTime: time.Now().UTC()}
	if err := ctx.ShouldBindQuery(&query); err != nil {
        ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
        return
    }

	restaurants, err := h.Service.GetRestaurantsByDate(query.DateTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of restaurants",
		"data": restaurants,
	})
}

func (h *Handler) GetRestaurantById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Restaurant not found",
		})
	}

	restaurant, err := h.Service.GetRestaurantById(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	if restaurant == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Restaurant not found",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Restaurant details",
		"data": restaurant,
	})
}

func (h *Handler) SearchRestaurant(ctx *gin.Context) {
	var query SearchQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
        ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
        return
    }

	restaurants, err := h.Service.SearchRestaurant(query.SearchQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of restaurants",
		"data": restaurants,
	})
}

func (h *Handler) SearchDish(ctx *gin.Context) {
	var query SearchQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
        ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
        return
    }

	dishes, err := h.Service.SearchDish(query.SearchQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of dishes",
		"data": dishes,
	})
}


