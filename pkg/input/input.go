package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadInts() []int {
	lines := strings.Split(ReadInput(), "\n")
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
