package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mallet-zn/gopportunities/schemas"
)

// @Basepath /api/v1

// @Summary Update Opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Request body"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	id := ctx.Query("id")
	if id == "" {
		sendErr(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("error validation erro: %v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Update Opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendErr(ctx, http.StatusBadRequest, "error updating opening")
		return
	}

	sendSucess(ctx, "update-opening", opening)
}
