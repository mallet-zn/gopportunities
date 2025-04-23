package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mallet-zn/gopportunities/schemas"
)

// @Basepath /api/v1

// @Summary Delete Opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Request body"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendErr(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendErr(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSucess(ctx, "delete-opening", opening)
}
