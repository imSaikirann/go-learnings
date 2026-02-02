package main

import "github.com/gin-gonic/gin"

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var notes = []Note{}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/notes", func(c *gin.Context) {
		var note Note

		if err := c.ShouldBindBodyWithJSON(&note); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		note.ID = len(notes) + 1
		notes = append(notes, note)

		c.JSON(200, note)
	})

	r.GET("/notesz", func(c *gin.Context) {
		c.JSON(200, notes)
	})

	r.Run(":8080")
}
