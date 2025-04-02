package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Errors(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		fmt.Printf("\n\nstatus: %d\n\n", c.Writer.Status())

		if len(c.Errors) == 0 {
			return
		}

		status := c.Writer.Status()
		if status < 300 {
			status = http.StatusInternalServerError
		}

		errs := make([]string, 0, len(c.Errors))

		for _, e := range c.Errors.Errors() {
			errs = append(errs, e)
		}

		logger.Printf("status: %d, first err: %s", status, c.Errors.Errors()[0])

		c.JSON(status, gin.H{
			"errors": errs,
		})
	}
}
