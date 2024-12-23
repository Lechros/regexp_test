package main

import (
	"encoding/json"
	"github.com/BurntSushi/rure-go"
	"github.com/GRbit/go-pcre"
	"github.com/wasilibs/go-re2"
	"slices"
	//"github.com/flier/gohs/hyperscan"
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

func StandardMatchAll(pattern string) []string {
	regex := regexp.MustCompile("(?i)" + pattern)
	result := make([]string, 0, 10)
	for id, name := range names {
		if regex.MatchString(name) {
			result = append(result, id)
		}
	}
	return result
}

func StandardConcatMatchAll(pattern string) []string {
	regex := regexp.MustCompile("(?im)" + "^.*?" + pattern + ".*?$")
	result := make([]string, 0, 10)
	matches := regex.FindAllStringIndex(concatString, -1)
	for _, match := range matches {
		index, found := slices.BinarySearch(concatStartIndex, match[0])
		if !found { // Matched middle of name, index will be 1 bigger
			index--
		}
		id := concatIds[index]
		if len(result) == 0 || result[len(result)-1] != id {
			result = append(result, id)
		}
	}
	return result
}

func RuReMatchAll(pattern string) []string {
	regex := rure.MustCompile("(?i)" + pattern)
	result := make([]string, 0, 10)
	for id, name := range names {
		if regex.IsMatch(name) {
			result = append(result, id)
		}
	}
	return result
}

func RuReConcatMatchAll(pattern string) []string {
	regex := rure.MustCompile("(?im)" + "^.*?" + pattern + ".*?$")
	result := make([]string, 0, 10)
	matches := regex.FindAll(concatString)
	for i, match := range matches {
		if i%2 != 0 {
			continue
		}
		index, found := slices.BinarySearch(concatStartIndex, match)
		if !found { // Matched middle of name, index will be 1 bigger
			index--
		}
		id := concatIds[index]
		if len(result) == 0 || result[len(result)-1] != id {
			result = append(result, id)
		}
	}
	return result
}

func PcreMatchAll(pattern string) []string {
	regex := pcre.MustCompile(pattern, pcre.CASELESS)
	result := make([]string, 0, 10)
	for id, name := range names {
		if regex.MatchStringWFlags(name, pcre.CASELESS) {
			result = append(result, id)
		}
	}
	return result
}

func Re2MatchAll(pattern string) []string {
	regex := re2.MustCompile(pattern)
	result := make([]string, 0, 10)
	for id, name := range names {
		if regex.MatchString(name) {
			result = append(result, id)
		}
	}
	return result
}

//func HyperScanMatchAll(pattern string) {
//	p := hyperscan.NewPattern(pattern, hyperscan.Caseless)
//	db, _ := hyperscan.NewBlockDatabase(p)
//	for _, name := range names {
//		db.MatchString(name)
//	}
//}

//func StandardFindGroups(pattern string) {
//	regex := regexp.MustCompile("(?i)" + pattern)
//	for _, name := range names {
//		regex.FindAllStringSubmatchIndex(name, -1)
//	}
//}
//
//func RuReFindGroups(pattern string) {
//	regex := rure.MustCompile("(?i)" + pattern)
//	captures := regex.NewCaptures()
//	for _, name := range names {
//		regex.Captures(captures, name)
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
	builder := strings.Builder{}
	for id, gearData := range gears {
		var gear map[string]interface{}
		if err := json.Unmarshal(gearData, &gear); err != nil {
			log.Fatal(err)
		}
		name := gear["name"].(string)
		names[id] = strings.ToLower(name)

		// concat
		concatStartIndex = append(concatStartIndex, builder.Len())
		concatIds = append(concatIds, id)
		builder.WriteString(name)
		builder.WriteRune('\n')
	}
	concatString = builder.String()
	concatString = concatString[:len(concatString)-1]
}

func main() {
}
