package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"freeagent"
)

func main() {
	accessToken := os.Args[1]
	timeslipID := os.Args[2]

	fa := &freeagent.FreeAgent{
		Endpoint:    freeagent.SandboxEndpoint,
		AccessToken: accessToken,
	}

	timeslip, err := fa.GetTimeslip(timeslipID)
	if err != nil {
		log.Fatal(err)
	}

	timeslip.URL = ""
	timeslip.Hours = "1"

	timeslip, err = fa.PostTimeslip(timeslip)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := json.MarshalIndent(timeslip, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
