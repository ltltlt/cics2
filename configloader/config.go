package configloader

import (
	"log"
	"os"
	"strings"
)

// Loader abstract the load procedure
type Loader interface {
	Load(file string) error
}

// Config is a two-level map
type Config map[string]map[string]string

// config is the default Config
var config Config

// GetByL1 get level 1 config items
func GetByL1(l1 string) map[string]string {
	return config[l1]
}

// Get accept level1 and level2 as argument, return configure item
func Get(l1, l2 string) string {
	return config[l1][l2]
}

// GetOne accept level1 and level2 as argument, return first configure item
func GetOne(l1, l2 string) string {
	return config[l1][l2]
}

var loaderMap = make(map[string]Loader)

func init() {
	loaderMap["json"] = &JSON{}
	// read config file name from env
	uri := os.Getenv("CICS_CONFIG")
	if len(uri) == 0 {
		uri = "json:./configs/default.json" // you deserve this shit
	}
	parts := strings.Split(uri, ":")
	loader, file := parts[0], parts[1]
	if err := loaderMap[loader].Load(file); err != nil {
		log.Fatalln("init config fail:", err)
	}

	if len(os.Getenv("CICS_TEST")) > 0 {
		log.Printf("config content: +%v", config)
	}
}
