package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseError struct {
	Index int
	Word  string
	Err   error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("pkg pasre: error parsing %q as int", e.Word)
}

func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg:%v", r)
			}
		}
	}()

	fileds := strings.Fields(input)
	numbers = fileds2numbers(fileds)
	return
}

func fileds2numbers(fileds []string) (numbers []int) {
	if len(fileds) == 0 {
		panic("no word to parse")
	}

	for idx, filed := range fileds {
		num, err := strconv.Atoi(filed)
		if err != nil {
			panic(&ParseError{idx, filed, err})
		}

		numbers = append(numbers, num)
	}

	return
}
