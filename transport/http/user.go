package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func (h *Handler) GetUsers(ctx *gin.Context) {
	users, err := h.Service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of user",
		"data": users,
	})
}

func (h *Handler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Restaurant not found",
		})
	}

	user, err := h.Service.GetUserById(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User details info",
		"data": user,
	})
}

