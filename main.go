package main

import (
	"encoding/json" // to parse json
	"fmt"
	"io"       // will allow basic read and write commands
	"log"      // for logging messages to the console
	"net/http" // will allow us to make a get call to the service
)

type data struct {
	Endpoint       string                   `json:'endpoint'`
	Quotes         []map[string]interface{} `json:'quotes'`
	Requested_time string                   `json:'requested_time'`
	Timestamp      int32                    `json:'timestamp'`
}

func main() {

	//currencies := "EUROUSD,GBPUSD"
	url := "https://marketdata.tradermade.com/api/v1/convert?api_key=ZZXcztC_G_Aqlxvz6L0A&from=EUR&to=GBP&amount=1000"

	resp, getErr := http.Get(url) // http get request
	/*note: http.Get(url) is a function call to the Get function in the http package.
	It takes a URL as an argument and returns two values: resp and getErr.*/
	if getErr != nil { // this is an error handler to catch any error that ocurred during the get process
		log.Fatal(getErr)
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
	fmt.Println("---------")

	data_obj := data{}
	jsonErr := json.Unmarshal(body, &data_obj)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("endpoint", data_obj.Endpoint, "requested time", data_obj.Requested_time, "timestamp", data_obj.Timestamp)
	fmt.Println("---------")
	for key, value := range data_obj.Quotes {
		fmt.Println(key)
		fmt.Println("---------")
		fmt.Println("symbol", value["base_currency"], value["quote_currency"], "bid", value["bid"], "ask", value["ask"],
			"mid", value["mid"])

	}

}
