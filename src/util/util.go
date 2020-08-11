package util

import (
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func ReadLines(fname string) []string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}
	return lines
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Contains(col []string, item string) bool {
	for _, v := range col {
		if v == item {
			return true
		}
	}
	return false
}

func GetKeys(input map[string][]string) []string {
	result := []string{}
	for k, _ := range input {
		result = append(result, k)
	}
	return result
}

func RemoveDuplicates(items []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for v := range items {
		if seen[items[v]] != true {
			seen[items[v]] = true
			result = append(result, items[v])
		}
	}
	return result
}

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func IsPrime(number int) bool {
	switch {
	case number < 2:
		return false
	case number == 2:
		return true
	case number%2 == 0:
		return false
	default:
		for i := 3; (i * i) <= number; i += 2 {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}
