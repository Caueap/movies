package main

import(
	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

var movies = []Movie{}

func CreateMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.BindJSON(&newMovie); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	c.JSON(200, newMovie)

}

func main() {
	r := gin.Default()
	r.POST("/movie", CreateMovie)
	r.Run(":8080")
	
}