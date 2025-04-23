package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mallet-zn/gopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendErr(ctx, http.StatusInternalServerError, "error listing opening")
		return
	}

	sendSucess(ctx, "list-openings", openings)
}
