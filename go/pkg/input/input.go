package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInput(day int) string {
	content, err := ioutil.ReadFile(fmt.Sprintf("input/day%02d.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadInts(day int) []int {
	lines := strings.Split(ReadInput(day), "\n")
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	return nums
}
