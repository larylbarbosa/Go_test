package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	The Health API endpoint. Returns fixed payload on success.
*/
func (endpt *Endpoint) Health() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: send a string? Would be faster...
		ctx.JSON(http.StatusOK, gin.H{
			"type": "success",
			// "svc":  os.Getenv("SERVICE_NAME"),
		})
	}
}
