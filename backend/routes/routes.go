package routes

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	Message string `json:"message"`
}

const badtempEndpoint = "https://www.kalmar.se/opendata/WaterQualityObservedFiware.json"

func SetupRoutes() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// routes
	r.GET("/hello", hello)
	r.GET("/badtemp", badtemp)

	r.Run(":5000")
}

func hello(c *gin.Context) {
	msg := Msg{"Hello World!"}
	c.IndentedJSON(http.StatusOK, msg)
}

func badtemp(c *gin.Context) {
	resp, err := http.Get(badtempEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	data := string(body)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, data)
}
