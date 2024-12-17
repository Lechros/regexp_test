package main

import (
	"encoding/json"
	"github.com/BurntSushi/rure-go"
	"github.com/GRbit/go-pcre"
	//"github.com/flier/gohs/hyperscan"
	"github.com/wasilibs/go-re2"
	"log"
	"os"
	"regexp"
	"strings"
)

var gears map[string]json.RawMessage
var names map[string]string

func StandardMatchAll(pattern string) {
	regex := regexp.MustCompile("(?i)" + pattern)
	for _, name := range names {
		regex.MatchString(name)
	}
}

func StandardFindGroups(pattern string) {
	regex := regexp.MustCompile("(?i)" + pattern)
	for _, name := range names {
		regex.FindAllStringSubmatchIndex(name, -1)
	}
}

func RuReMatchAll(pattern string) {
	regex := rure.MustCompile("(?i)" + pattern)
	for _, name := range names {
		regex.IsMatch(name)
	}
}

func RuReFindGroups(pattern string) {
	regex := rure.MustCompile("(?i)" + pattern)
	captures := regex.NewCaptures()
	for _, name := range names {
		regex.Captures(captures, name)
	}
}

func PcreMatchAll(pattern string) {
	regex := pcre.MustCompile(pattern, pcre.CASELESS)
	for _, name := range names {
		regex.MatchStringWFlags(name, pcre.CASELESS)
	}
}

func Re2MatchAll(pattern string) {
	regex := re2.MustCompile(pattern)
	for _, name := range names {
		regex.MatchString(name)
	}
}

//func HyperScanMatchAll(pattern string) {
//	p := hyperscan.NewPattern(pattern, hyperscan.Caseless)
//	db, _ := hyperscan.NewBlockDatabase(p)
//	for _, name := range names {
//		db.MatchString(name)
//	}
//}

func init() {
	data, err := os.ReadFile("gear-data.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &gears); err != nil {
		log.Fatal(err)
	}

	names = make(map[string]string, len(gears))
	for id, gearData := range gears {
		var gear map[string]interface{}
		if err := json.Unmarshal(gearData, &gear); err != nil {
			log.Fatal(err)
		}
		names[id] = strings.ToLower(gear["name"].(string))
	}
}

func main() {
}
