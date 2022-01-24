package validate

import (
	"strconv"
	"strings"
)

func BusinessNumber(bisNo string) bool {
	bisNo = strings.Replace(bisNo, "-", "", 2)
	var arrCheckNum = []int{1, 3, 7, 1, 3, 7, 1, 3, 5}

	sum := 0
	if len(bisNo) == 10 {
		for i := 0; i < 9; i++ {
			r, _ := strconv.Atoi(string(bisNo[i]))
			sum += r * arrCheckNum[i]
		}
		t, _ := strconv.Atoi(string(bisNo[8]))
		sum += (t * arrCheckNum[8]) / 10
		sum = 10 - (sum % 10)

		lastNum, _ := strconv.Atoi(string(bisNo[9]))

		if sum == lastNum {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}
