package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/bayugyug/rest-building/config"
	"github.com/bayugyug/rest-building/controllers"
)

const (
	//VersionMajor main ver no.
	VersionMajor = "0.1"
	//VersionMinor sub  ver no.
	VersionMinor = "0"
)

var (
	//BuildTime pass during build time
	BuildTime string
	//APIVersion is the app ver string
	APIVersion string
)

//internal system initialize
func init() {
	//uniqueness
	rand.Seed(time.Now().UnixNano())
	APIVersion = "Ver: " + VersionMajor + "." + VersionMinor + "-" + BuildTime

}

func main() {

	start := time.Now()
	log.Println(APIVersion)

	var err error

	//init
	appcfg := config.NewAppSettings()

	//check
	if appcfg.Config == nil {
		log.Fatal("Oops! Config missing")
	}

	//init service
	if controllers.APIInstance, err = controllers.NewAPIService(
		controllers.WithSvcOptAddress(":" + appcfg.Config.Port),
	); err != nil {
		log.Fatal("Oops! config might be missing", err)
	}

	//run service
	controllers.APIInstance.Run()
	log.Println("Since", time.Since(start))
	log.Println("Done")
}
