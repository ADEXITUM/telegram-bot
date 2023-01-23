package checker

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func LoadConfig(filename string) ([]string, error) {
	config, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	c := make(map[string]json.RawMessage)

	err = json.Unmarshal(config, &c)
	if err != nil {
		fmt.Println(err)
	}

	var links = []string{}

	for _, v := range c {
		links = append(links, string(v))
	}

	return links, nil
}

func extractValue(body string, key string) string {
	keystr := "\"" + key + "\":[^,;\\]}]*"
	r, _ := regexp.Compile(keystr)
	match := r.FindString(body)
	keyValMatch := strings.Split(match, ":")
	return strings.ReplaceAll(keyValMatch[1], "\"", "")
}

func GetLinks() []string {
	links, err := LoadConfig("config.json")
	if err != nil {
		fmt.Println("can't get links")
	}
	return links
}
