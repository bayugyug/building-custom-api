package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/bayugyug/building-custom-api/api/routes"
	"github.com/bayugyug/building-custom-api/configs"
)

const (
	// VersionMajor main ver no.
	VersionMajor = "0.1"
	// VersionMinor sub  ver no.
	VersionMinor = "0"
)

var (
	// BuildTime pass during build time
	BuildTime string
	// APIVersion is the app ver string
	APIVersion string
)

//init internal system initialize
func init() {
	//uniqueness
	rand.Seed(time.Now().UnixNano())
	APIVersion = "Ver: " + VersionMajor + "." + VersionMinor + "-" + BuildTime

}

func main() {
	var err error
	start := time.Now()
	log.Println(APIVersion)
	//init
	appcfg := configs.NewAppSettings()
	//check
	if appcfg.Config == nil {
		log.Fatal("Oops! Config missing")
	}
	//init service
	if routes.Service, err = routes.NewAPIService(
		routes.WithSvcOptAddress(":" + appcfg.Config.Port),
	); err != nil {
		log.Fatal("Oops! config might be missing", err)
	}
	//run service
	routes.Service.Run()
	log.Println("Since", time.Since(start))
	log.Println("Done")
}
