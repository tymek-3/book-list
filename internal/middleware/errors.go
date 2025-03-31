package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO:

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

		type err struct {
			msg   string
			error string
		}
		errs := make([]err, 0)

		for _, e := range c.Errors.Errors() {
			errs = append(errs, err{"error", e})
		}

		logger.Printf("status: %d, first err: %s", status, c.Errors.Errors()[0])

		c.JSON(status, gin.H{
			"message": errs,
		})
	}
}
