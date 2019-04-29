package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/bayugyug/rest-building/utils"
)

const (
	//status
	usageConfig = "use to set the config file parameter with HTTP-port"
)

var (
	//Settings of the app
	Settings *APISettings
)

// ParameterConfig optional parameter structure
type ParameterConfig struct {
	Port    string `json:"port"`
	Showlog bool   `json:"showlog"`
}

// APISettings is a config mapping
type APISettings struct {
	Config    *ParameterConfig
	CmdParams string
	EnvVars   map[string]*string
}

// Setup options settings
type Setup func(*APISettings)

// WithSetupConfig for cfg
func WithSetupConfig(r *ParameterConfig) Setup {
	return func(args *APISettings) {
		args.Config = r
	}
}

// WithSetupCmdParams for the json params
func WithSetupCmdParams(r string) Setup {
	return func(args *APISettings) {
		args.CmdParams = r
	}
}

// WithSetupEnvVars for envt params
func WithSetupEnvVars(r map[string]*string) Setup {
	return func(args *APISettings) {
		args.EnvVars = r
	}
}

// NewAppSettings main entry for config
func NewAppSettings(setters ...Setup) *APISettings {
	//set default
	cfg := &APISettings{
		EnvVars: make(map[string]*string),
	}
	//maybe export from envt
	cfg.EnvVars = map[string]*string{
		"API_CONFIG": &cfg.CmdParams,
	}
	//chk the passed params
	for _, setter := range setters {
		setter(cfg)
	}
	//start
	cfg.Initializer()
	return cfg
}

//InitRecov is for dumpIng segv in
func (g *APISettings) InitRecov() {
	//might help u
	defer func() {
		recvr := recover()
		if recvr != nil {
			log.Println("MAIN-RECOV-INIT: ", recvr)
		}
	}()
}

//InitEnvParams enable all OS envt vars to reload internally
func (g *APISettings) InitEnvParams() {
	//just in-case, over-write from ENV
	for k, v := range g.EnvVars {
		if os.Getenv(k) != "" {
			*v = os.Getenv(k)
		}
	}
	//get options
	flag.StringVar(&g.CmdParams, "config", g.CmdParams, usageConfig)
	flag.Parse()
}

//Initializer set defaults for initial reqmts
func (g *APISettings) Initializer() {
	//prepare
	g.InitRecov()
	g.InitEnvParams()
	log.Println("CmdParams:", g.CmdParams)

	//try to reconfigure if there is passed params, otherwise use show err
	if g.CmdParams != "" {
		g.Config = g.FormatParameterConfig(g.CmdParams)
	}

	//check defaults
	if g.Config == nil {
		return
	}
	//set dump flag
	utils.ShowMeLog = g.Config.Showlog

}

//FormatParameterConfig new ParameterConfig
func (g *APISettings) FormatParameterConfig(s string) *ParameterConfig {
	var cfg ParameterConfig
	if err := json.Unmarshal([]byte(s), &cfg); err != nil {
		log.Println("FormatParameterConfig", err)
		return nil
	}
	return &cfg
}
