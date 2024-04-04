package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseBody struct {
	DiashId		int		`json:"dishId" binding:"required"`
	UserId		int		`json:"userId" binding:"required"`
}

func (h *Handler) Purchase(ctx *gin.Context) {
	var body PurchaseBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
        return
    }

	if err := h.Service.Purchase(body.UserId, body.DiashId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Purchase is successfull",
	})
}

