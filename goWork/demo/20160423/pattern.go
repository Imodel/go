package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	serchInt := "jack:123.23 lucky:44.34 ketty:90.5"

	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(serchInt)); ok {
		fmt.Println("match found")
	}

	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(serchInt, "##.#")

	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(serchInt, f)

	fmt.Println(str2)
}
