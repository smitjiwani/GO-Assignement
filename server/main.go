package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var numbers []int

func setupRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	
	api.POST("/numbers", handlePostNumber)
	api.GET("/numbers", handleGetNumbers)
	
	return router
}

func main() {
	router := setupRoutes()
	router.Run(":8080")
}

func handlePostNumber(c *gin.Context) {
	type request struct {
		Number int `json:"number"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Number == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Zero is not allowed"})
		return
	}

	if len(numbers) == 0 {
		numbers = append(numbers, req.Number)
	} else if len(numbers) == 1 {
		if (numbers[0] > 0 && req.Number > 0) || (numbers[0] < 0 && req.Number < 0) {
			numbers = append(numbers, req.Number)
		} else {
			remainingValue := abs(req.Number)
			if abs(numbers[0]) <= remainingValue {
				remainingValue -= abs(numbers[0])
				numbers = numbers[:0]
				if remainingValue > 0 {
					if req.Number > 0 {
						numbers = append(numbers, remainingValue)
					} else {
						numbers = append(numbers, -remainingValue)
					}
				}
			} else {
				if numbers[0] > 0 {
					numbers[0] -= remainingValue
				} else {
					numbers[0] += remainingValue
				}
			}
		}
	} else {
		if (numbers[0] >= 0 && req.Number >= 0) || (numbers[0] < 0 && req.Number < 0) {
			numbers = append(numbers, req.Number)
		} else {
			remainingValue := abs(req.Number)
			for i := 0; i < len(numbers) && remainingValue > 0; {
				if abs(numbers[i]) <= remainingValue {
					remainingValue -= abs(numbers[i])
					numbers = append(numbers[:i], numbers[i+1:]...)
				} else {
					if numbers[i] > 0 {
						numbers[i] -= remainingValue
					} else {
						numbers[i] += remainingValue
					}
					remainingValue = 0
				}
			}

			print(remainingValue)
			if remainingValue > 0 {
				if req.Number < 0 {
					numbers = append(numbers, -remainingValue)
				} else {
					numbers = append(numbers, remainingValue)
				}
			}
		}
	}
	
	c.JSON(http.StatusOK, gin.H{"list": numbers})
}

func handleGetNumbers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": numbers})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}