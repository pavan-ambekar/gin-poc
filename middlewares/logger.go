package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setupLofOutput() {
	file, error := os.Create("gin.log")
	if error != nil {
		panic("failed to create log file.")
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func Logger() gin.HandlerFunc {
	setupLofOutput()
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		// return params.Method + " " + params.Path
		return fmt.Sprintf("%s - [%v] %-7s %#v %d %v \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
