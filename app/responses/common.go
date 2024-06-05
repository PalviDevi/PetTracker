package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessApiResponse(c *gin.Context, transKey string, transParams map[string]string, data interface{}) {
	// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
	c.PureJSON(http.StatusOK,
		map[string]interface{}{
			"status":  1,
			"message": transKey,
			"data":    data,
		})
}
