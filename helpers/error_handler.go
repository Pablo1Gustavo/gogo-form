package helpers

import (
	"gogo-form/domain"

	"github.com/gin-gonic/gin"
)

func RespondToError(ctx *gin.Context, err error) bool {
	requestError, ok := err.(*domain.RequestError)
	if !ok {
		return false
	}

	var obj gin.H

	switch requestError.Code {
	case 0:
		requestError.Code = 500
		obj = gin.H{"message": "Unexpected server error"}
	case 404:
		obj = gin.H{"message": "Resource not found"}
	case 422:
		obj = gin.H{"message": "Invalid structure", "errors": requestError.Details}
	}

	ctx.JSON(requestError.Code, obj)

	return true
}
