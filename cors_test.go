package gin_cors

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCors_CorsMiddleware(t *testing.T) {

	r := gin.Default()

	cors := &Cors{
		AccessControlAllowOrigins: []string{
			"http://localhost",
		},
		AccessControlAllowMethods: "PUT,POST,GET,DELETE",
	}

	r.Use(cors.CorsMiddleware())

	r.POST("/te", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "this is data",
		})
	})

	r.Run(":7878")

}
