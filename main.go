package main

import (
	"context"
  "encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type jsonResp struct {
  ID    int `json:"id"`
  Time  string `json:"time"`
  Words []string `json:"words"`
}

func main() {
	var in, fTime, tTime, wWord string
	var wID int

	input := flag.String("input", "", "An S3 URI (`s3://{bucket}/{key}`) that refers to the source object to be filtered.")
	withID := flag.Int("with-id", 0, "An integer that contains the `id` of a JSON object to be selected.")
	fromTime := flag.String("from-time", "", "An RFC3339 timestamp that represents the earliest `time` of a JSON object to be selected.")
	toTime := flag.String("to-time", "", "An RFC3339 timestamp that represents the latest `time` of JSON object to be selected.")
	withWord := flag.String("with-word", "", "A string containing a word that must be contained in `words` of a JSON objec to be selected.")

	flag.Parse()

	// Check flags
	if *input == "" {
		fmt.Println("input (-input) flag required")
		os.Exit(1)
    
	} else {
		fmt.Printf("found input (-input) flag: %s\n", *input)
		in = *input
		fmt.Println(in)
	}

	if *withID != 0 {
		fmt.Printf("found withID (-with-id) flag: %d\n", *withID)
		wID = *withID
		fmt.Println(wID)
	}

	if *fromTime != "" {
		fmt.Printf("found fromTime (-from-time) flag: %s\n", *fromTime)
		fTime = *fromTime
		fmt.Println(fTime)
	}

	if *toTime != "" {
		fmt.Printf("found toTime (-to-time) flag: %s\n", *toTime)
		tTime = *toTime
		fmt.Println(tTime)
	}

	if *withWord != "" {
		fmt.Printf("found withWord (-with-word) flag: %s\n", *withWord)
		wWord = *withWord
		fmt.Println(wWord)
	}

}
