package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"ward-go/common"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v\n", r)
				debug.PrintStack()
				c.HTML(http.StatusInternalServerError, "error/500.gohtml", gin.H{"theme": common.Config.Theme})
			}
		}()
		c.Next()

	}
}
