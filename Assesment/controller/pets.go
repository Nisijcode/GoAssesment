package controller

import (
	"github.com/gin-gonic/gin"
)

const (
	// Success Status: Success
	Success = 0
	// Failed Status: Failed
	Failed = 1
	// NotFound Status: Not Found
	NotFound = 2
	// Unauthorised Staus: Unauthorised
	Unauthorised = 3
	// DatabaseConnectionError Status: Database connection error
	DatabaseConnectionError = 4
	// DatabaseError Status: Database error
	DatabaseError = 5
	// InternalServerError Status: Internal server error
	InternalServerError = 6
	// BadRequestError Status: Bad Request error
	BadRequestError = 7
)

var responseCodes = []struct {
	Status int
	Result string
}{
	{Status: 200, Result: "Success"},
	{Status: 200, Result: "Failed"},
	{Status: 404, Result: "NotFound"},
	{Status: 401, Result: "UnAuthorised"},
	{Status: 500, Result: "DatabaseConnectionError"},
	{Status: 500, Result: "DatabaseError"},
	{Status: 500, Result: "InternalServerError"},
	{Status: 400, Result: "BadRequestError"},
}

//Response Format of the response
func Response(result int, message string, data interface{}) gin.H {
	if data != nil {
		return gin.H{
			"status":  responseCodes[result].Status,
			"result":  responseCodes[result].Result,
			"message": message,
			"data":    data,
		}
	}
	return gin.H{
		"status":  responseCodes[result].Status,
		"result":  responseCodes[result].Result,
		"message": message,
	}
}
