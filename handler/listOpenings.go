package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mallet-zn/gopportunities/schemas"
)

// @Basepath /api/v1

// @Summary List Openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendErr(ctx, http.StatusInternalServerError, "error listing opening")
		return
	}

	sendSucess(ctx, "list-openings", openings)
}
