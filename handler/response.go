package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mallet-zn/gopportunities/schemas"
)

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode string `json:"statusCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ListOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

func sendErr(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":    msg,
		"statusCode": code,
	})
}

func sendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data":    data,
	})
}
