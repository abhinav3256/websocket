package main

import (
	"github.com/gin-gonic/gin"
)

type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
	H string `json:"h"`
}

func main() {
	//getQuote()
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// //http.HandleFunc("/ws", registerClient)

}

func setupRoutes(r *gin.Engine) {

	r.GET("wp", registerClient)

}
