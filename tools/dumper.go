package tools

import (
	"encoding/json"
	"log"
)

// ShowMeLog flag to switch verbose
var ShowMeLog = true

// Dumper verbose logs
func Dumper(infos ...interface{}) {
	if ShowMeLog {
		j, _ := json.MarshalIndent(infos, "", "\t")
		log.Println(string(j))
	}
}
