package main

import (
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// prepare signal nitification channel
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// watch for notification events and clean up on exit
	go func() {
		for {
			switch <-signalChan {
			case syscall.SIGHUP:
				os.Exit(int(syscall.SIGHUP))
			case syscall.SIGINT:
				os.Exit(int(syscall.SIGINT))
			case syscall.SIGTERM:
				os.Exit(int(syscall.SIGTERM))
			case syscall.SIGQUIT:
				os.Exit(int(syscall.SIGQUIT))
			}
		}
	}()

	// set release mode
	gin.SetMode(gin.ReleaseMode)

	// create a gin router with default middleware
	r := gin.Default()

	// default API func
	f := func(c *gin.Context) {
		var in In

		// get aloritm from registered route
		algo := strings.TrimPrefix(c.FullPath(), "/")
		if !(algo == BaseTypeAlgo || algo == Custom1TypeAlgo || algo == Custom2TypeAlgo) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "wrong algoritm"}) // should never happen

			return
		}

		// decode input data to object
		if err := c.ShouldBind(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		// get output data
		out, err := in.Process(algo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, out)
	}

	// register Base API endpoint
	/*
	   curl --http1.1 -XPOST -H "Content-type: application/json" -H "Accept: application/json" -d '{
	     "A": false,
	     "B": true,
	     "C": true,
	     "D": 10.0,
	     "E": 50,
	     "F": 150
	   }' 'http://127.0.0.1:8080/base'
	*/
	r.POST("/"+BaseTypeAlgo, f)

	// register Custom1 API endpoint
	/*
	   curl --http1.1 -XPOST -H "Content-type: application/json" -H "Accept: application/json" -d '{
	     "A": true,
	     "B": true,
	     "C": false,
	     "D": 5.0,
	     "E": 10,
	     "F": 20
	   }' 'http://127.0.0.1:8080/custom1'
	*/
	r.POST("/"+Custom1TypeAlgo, f)

	// register Custom2 API endpoint
	/*
	   curl --http1.1 -XPOST -H "Content-type: application/json" -H "Accept: application/json" -d '{
	     "A": true,
	     "B": true,
	     "C": true,
	     "D": 55.5,
	     "E": 210,
	     "F": 40
	   }' 'http://127.0.0.1:8080/custom2'
	*/
	r.POST("/"+Custom2TypeAlgo, f)

	// listen and serve API daemon
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
