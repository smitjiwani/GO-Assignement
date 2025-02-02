package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"os"
)

var numbers []int

func main() {
	router := gin.Default()
	api := router.Group("/api")

	api.POST("/numbers", func(c *gin.Context) {
		type request struct {
			Number int `json:"number"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(numbers) == 0 {
			if req.Number < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Negative numbers not allowed if the array is empty"})
				return
			}
			numbers = append(numbers, req.Number)
		} else if (numbers[0] >= 0 && req.Number >= 0) || (numbers[0] < 0 && req.Number < 0) {
			numbers = append(numbers, req.Number)
		} else {
			removeQty := req.Number
			if removeQty < 0 {
				removeQty = -removeQty
			}
			for removeQty > 0 && len(numbers) > 0 {
				i := len(numbers) - 1
				if removeQty >= abs(numbers[i]) {
					removeQty -= abs(numbers[i])
					numbers = numbers[:i]
				} else {
					if numbers[i] > 0 {
						numbers[i] -= removeQty
					} else {
						numbers[i] += removeQty
					}
					removeQty = 0
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"list": numbers})
	})

	api.GET("/numbers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"list": numbers})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	router.Run(":" + port)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}