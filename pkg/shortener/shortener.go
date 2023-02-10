package shortener

import (
	"fmt"
	"strings"
)

const (
	alphabet    = "vuwev84g3hENYerhZXCQWE44renfdVWnyj9349AMOGUSgTME657W7vnxd75ev297v24KFNEie57tje56856t25t2owgetbimvwi67mvwA"
	alphabetLen = uint32(len(alphabet))
	BaseURL     = "http://localhost:8080/"
)

func GetShortURL(linkID uint32) string {
	var nums []uint32

	for linkID > 0 {
		nums = append(nums, linkID%alphabetLen)
		linkID = linkID / alphabetLen
	}

	nums = Reverse(nums)

	var chars []string
	for _, num := range nums {
		chars = append(chars, string(alphabet[num]))
	}

	return fmt.Sprintf("%s%s", BaseURL, strings.Join(chars, ""))
}

func Reverse(nums []uint32) []uint32 {
	for i := 0; i < len(nums); i++ {
		nums = append(nums[1:], nums[0])
	}
	return nums
}
