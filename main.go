package main

import (
	"encoding/json"
	"github.com/BurntSushi/rure-go"
	"github.com/GRbit/go-pcre"
	"slices"
	//"github.com/flier/gohs/hyperscan"
	"github.com/wasilibs/go-re2"
	"log"
	"os"
	"regexp"
	"strings"
)

var gears map[string]json.RawMessage
var names map[string]string
var concatString string
var concatStartIndex []int
var concatIds []string

func StandardMatchAll(pattern string) {
	regex := regexp.MustCompile("(?i)" + pattern)
	for _, name := range names {
		regex.MatchString(name)
	}
}

func StandardConcatMatchAll(pattern string) {
	regex := regexp.MustCompile("(?i)" + pattern)
	matches := regex.FindAllStringIndex(concatString, -1)
	for _, match := range matches {
		index, found := slices.BinarySearch(concatStartIndex, match[0])
		if found { // Matched start of name
			_ = concatIds[index]
		} else { // Matched middle of name, index will be 1 bigger
			_ = concatIds[index-1]
		}
	}
}

func RuReMatchAll(pattern string) {
	regex := rure.MustCompile("(?i)" + pattern)
	for _, name := range names {
		regex.IsMatch(name)
	}
}

func RuReConcatMatchAll(pattern string) {
	regex := rure.MustCompile("(?i)" + pattern)
	matches := regex.FindAll(concatString)
	for i, match := range matches {
		if i%2 != 0 {
			continue
		}
		index, found := slices.BinarySearch(concatStartIndex, match)
		if found { // Matched start of name
			_ = concatIds[index]
		} else { // Matched middle of name, index will be 1 bigger
			_ = concatIds[index-1]
		}
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

func StandardFindGroups(pattern string) {
	regex := regexp.MustCompile("(?i)" + pattern)
	for _, name := range names {
		regex.FindAllStringSubmatchIndex(name, -1)
	}
}

func RuReFindGroups(pattern string) {
	regex := rure.MustCompile("(?i)" + pattern)
	captures := regex.NewCaptures()
	for _, name := range names {
		regex.Captures(captures, name)
	}
}

func init() {
	data, err := os.ReadFile("gear-data.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &gears); err != nil {
		log.Fatal(err)
	}

	names = make(map[string]string, len(gears))
	builder := strings.Builder{}
	for id, gearData := range gears {
		var gear map[string]interface{}
		if err := json.Unmarshal(gearData, &gear); err != nil {
			log.Fatal(err)
		}
		name := gear["name"].(string)
		names[id] = strings.ToLower(name)

		// concat
		builder.WriteString(name)
		builder.WriteRune('\n')
		concatStartIndex = append(concatStartIndex, builder.Len())
		concatIds = append(concatIds, id)
	}
	concatString = builder.String()
	concatString = concatString[:len(concatString)-1]
}

func main() {
}
