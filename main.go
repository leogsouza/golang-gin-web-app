package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Joke contains information about a single joke
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding: "required"`
}

// List of jokes
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

func main() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	// Our api will consist of just two routes
	// /jokes - which will retrieve a list of jokes a user can see
	// /jokes/like/:jokeID - wich wull capture likes sent to a particular joke
	api.GET("/jokes", JokeHandler)
	api.POST("/jokes/like/:jokeID", LikeHandler)

	// Start and run the server
	router.Run(":3000")
}

// JokeHandler retrieves a list of available joke
func JokeHandler(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, jokes)
}

// LikeHandler increments the likes of a particular joke
func LikeHandler(c *gin.Context) {
	// confirm is Joke ID sent is valid
	// remember to import `strconv`package
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		// find jokes and increment likes
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes++
			}
		}

		// return a pointer to hte updated jokes list
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
