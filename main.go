package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func message(text string, url string){
	text = fmt.Sprintf(`{"text":"%s"}`, text)
	var jsonStr = []byte(text)
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	
	req.Header.Add("Content-Type","application/json")
	client := &http.Client{}
	
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}	
}

func main()  {
	err := godotenv.Load()
  if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	url := os.Getenv("GCHAT_WH")
	port := os.Getenv("GCHAT_PORT")

	r := gin.Default()

	type RecognitionRequestBody struct {
		Person string `json:"person"`
		Reason string `json:"reason"`
	}

	r.POST("/recognize", func(c *gin.Context) {
		var body RecognitionRequestBody
		if err := c.BindJSON(&body); err != nil {
				fmt.Println(err)
		}
		message(fmt.Sprintf(`Recognition to %s! %s`, body.Person, body.Reason),url)
  })

	r.Run(fmt.Sprintf(":%s",port))
}
