package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Temp struct {
	Value string `json:"value"`
}

type Date struct {
	Type  string `json:"-"`
	Value string `json:"value"`
}

type Beach struct {
	Name     string `json:"id"`
	Type     string `json:"-"`
	Date     Date   `json:"dateObserved"`
	Temp     Temp   `json:"temperature"`
	Location string `json:"-"`
}

type FormatedTemp struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Temperature string `json:"temperature"`
}

const badtempEndpoint = "https://www.kalmar.se/opendata/WaterQualityObservedFiware.json"

func SetupRoutes() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// routes
	r.GET("/api/hello", hello)
	r.GET("/api/badtemp", badtemp)

	r.Run(":5000")
}

func hello(c *gin.Context) {
	msg := "Hello World!"
	c.IndentedJSON(http.StatusOK, msg)
}

func badtemp(c *gin.Context) {
	// GET JSON data
	resp, err := http.Get(badtempEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Read data
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Stringify data
	dataString := string(body)

	// Unmarshal into an array
	var beachTemps []Beach

	err = json.Unmarshal([]byte(dataString), &beachTemps)

	if err != nil {
		log.Fatal(err)
	}

	// format data
	var formatedTemps []FormatedTemp

	for i, e := range beachTemps {
		// format date
		date := e.Date.Value
		dateTime := strings.Split(date, "T")
		date = dateTime[0]  // the date part
		time := dateTime[1] // the time part
		timeArr := strings.Split(time, ":")
		time = timeArr[0] + ":" + timeArr[1]

		// format name
		name := e.Name
		name = strings.Split(name, ":")[2]

		// Round to whole number
		tempString := e.Temp.Value
		res, err := strconv.ParseFloat(tempString, 32)
		temp := math.Round(res)

		if err != nil {
			log.Fatal(err)
		}

		// add to array
		formatedTemp := FormatedTemp{
			Id:          fmt.Sprint(i),
			Name:        name,
			Date:        date,
			Time:        time,
			Temperature: fmt.Sprint(temp),
		}
		formatedTemps = append(formatedTemps, formatedTemp)
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, formatedTemps)
}
