package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MedicalRecord struct{}

func main() {
	var record MedicalRecord

	err := json.NewDecoder(os.Stdin).Decode(&record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(record)
}
