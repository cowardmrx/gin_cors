package gin_cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cors struct {
	AccessControlAllowOrigins     []string
	AccessControlAllowHeaders     string
	AccessControlAllowMethods     string
	AccessControlExposeHeaders    string
	AccessControlAllowCredentials string
}

func (cors *Cors) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range cors.AccessControlAllowOrigins {
			if v == c.Request.Header.Get("Origin") {
				c.Writer.Header().Set("Access-Control-Allow-Origin", v)
				c.Writer.Header().Set("Access-Control-Allow-Headers", cors.AccessControlAllowHeaders)
				c.Writer.Header().Set("Access-Control-Allow-Methods", cors.AccessControlAllowMethods)
				c.Writer.Header().Set("Access-Control-Expose-Headers", cors.AccessControlExposeHeaders)
				c.Writer.Header().Set("Access-Control-Allow-Credentials", cors.AccessControlAllowCredentials)
			}
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
