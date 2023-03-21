package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func message(text string, url string){
	text := fmt.Sprintf(`{"text":"%s"}`, text)
	var jsonStr = []byte(`{"text":"Hello World."}`)
	url := ""
	
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

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

    c.JSON(http.StatusOK, gin.H{
      "message": "bingbong",
    })
  })

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
