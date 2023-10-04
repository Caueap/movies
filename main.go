package main

import (
	"github.com/gin-gonic/gin"
)

// Movie represents a movie entity.
type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var movies = []Movie{}

// CreateMovie creates a new movie entry.
// @Summary Create a new movie
// @Description Creates a new movie entry.
// @Accept json
// @Produce json
// @Param movie body Movie true "Movie data"
// @Success 201 {object} Movie
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /movie [post]
func CreateMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.BindJSON(&newMovie); err != nil {
		c.JSON(400, ErrorResponse{Error: "Bad request"})
		return
	}

	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	c.JSON(201, newMovie)
}

// GetMovies returns all registered movies.
// @Summary Get all movies
// @Description Get all the registered movies
// @Produce json
// @Success 200 {array} Movie
// @Router /movies [get]
func GetMovies(c *gin.Context) {
	c.JSON(200, movies)
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	r := gin.Default()
	r.POST("/movie", CreateMovie)
	r.GET("/movies", GetMovies)
	r.Run(":8080")
}