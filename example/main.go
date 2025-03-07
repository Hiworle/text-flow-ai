package main

import (
	"flow/constant"
	"flow/openai"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.POST("/process", func(c *gin.Context) {
		var req constant.Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		// Change to your own openai api key
		flow := openai.New(openai.WithBaseUrl("https://api.openai.com/v1"),
			openai.WithToken(os.Getenv("OPENAI_API_KEY")),
			openai.WithModel("gpt-4"))
		flow.Init(constant.MakeSystemPrompt(req.Examples, req.Rules))

		output, err := flow.Process(c, req.Input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": output})
	})

	r.Run(":8000")
	return
}
