package middleware

import (
	"net/http"

	"github.com/porrporporrpor/covid-summary/model"

	"github.com/gin-gonic/gin"
)

func CustomRecovery(c *gin.Context, recovered interface{}) {
	err, ok := recovered.(string)
	if ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.HttpResponse{
			Status: model.PanicStatus,
			Data:   err,
		})
	}
}
