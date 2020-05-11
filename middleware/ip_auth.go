package middleware

import (
	"fmt"
	"tbwisk/public"

	"github.com/gin-gonic/gin"
)

//IPAuthMiddleware ip网关
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range public.GetStringSliceConf("base", "allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			ResponseError(c, InternalErrorCode, fmt.Errorf("%v, not in iplist", c.ClientIP()))
			c.Abort()
			return
		}
		c.Next()
	}
}
