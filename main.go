package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tushargarg0987/amazon-scraper/helper"
	"github.com/tushargarg0987/amazon-scraper/scraper"
)

func scrapeRequestHandler(req *gin.Context) {
	searchQuery := req.Param("query")
	if searchQuery == "" {
		req.IndentedJSON(http.StatusPartialContent, gin.H{"error": "Expected some query parameters"})
	}
	newSeachQuery := helper.QueryAdjuster(searchQuery)
	data, err := scraper.Scrap(newSeachQuery)
	if err != nil {
		fmt.Println("Check network connection and try again")
	}
	var result []map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	req.IndentedJSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, welcome to the amazon web scraping service")
	})
	router.GET("/scrape/:query", scrapeRequestHandler)
	router.Run("localhost:8000")
}
