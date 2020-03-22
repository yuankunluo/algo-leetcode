package main

import "regexp"

func checkRecord(s string) bool {
	reg, _ := regexp.Compile("A.*A|LLL")
	return !reg.MatchString(s)
}
