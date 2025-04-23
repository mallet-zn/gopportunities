package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mallet-zn/gopportunities/schemas"
)

// @Basepath /api/v1

// @Summary Show Job Opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Request body"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendErr(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSucess(ctx, "show-opening", opening)
}
