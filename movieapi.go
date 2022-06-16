package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/movies/year/:year", year)
	r.GET("/movies/rating/:rating", rating)
	r.GET("/movies/genre/:genre", genre)
}

//Dummy function

func year(c *gin.Context) {
	year, ok := c.Params.Get("year")
	records := readCsvFile("./movies.csv")
	moviename := newFunction1(records, year)

	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{
		"year":  year,
		"movie": moviename,
	}
	c.JSON(http.StatusOK, res)
}

func rating(c *gin.Context) {
	rating, ok := c.Params.Get("rating")
	records := readCsvFile("./movies.csv")
	moviename := newFunction2(records, rating)

	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"rating": rating,
		"movie":  moviename,
	}
	c.JSON(http.StatusOK, res)

}

func genre(c *gin.Context) {
	genre, ok := c.Params.Get("genre")
	records := readCsvFile("./movies.csv")
	moviename := newFunction3(records, genre)

	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"genre": genre,
		"movie": moviename,
	}
	c.JSON(http.StatusOK, res)
}

func newFunction1(records [][]string, year string) []string {
	var moviename string
	var moviearr = []string{}
	for i := 1; i < len(records); i++ {

		if records[i][7] == year {
			//	fmt.Println(records[i][j+1])
			moviename = records[i][0]
			moviearr = append(moviearr, moviename)

		}
	}
	return moviearr
}

func newFunction2(records [][]string, rating string) []string {
	var moviename string
	var moviearr = []string{}
	for i := 1; i < len(records); i++ {

		if records[i][5] >= rating {
			//	fmt.Println(records[i][j+1])
			moviename = records[i][0]
			moviearr = append(moviearr, moviename)

		}
	}
	return moviearr
}

func newFunction3(records [][]string, genre string) []string {
	var moviename string
	var moviearr = []string{}
	for i := 1; i < len(records); i++ {

		if records[i][1] == genre {
			//	fmt.Println(records[i][j+1])
			moviename = records[i][0]
			moviearr = append(moviearr, moviename)

		}
	}
	return moviearr
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
