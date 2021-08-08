package demo_leetcode

import (
	"fmt"
	"testing"
)

func TestStrToTwoDimensionIntegerArray(t *testing.T) {
	str := "[[2,1,1],[1,1,0],[0,1,1]]"
	integerArray := strToTwoDimensionIntegerArray(str)
	fmt.Println(integerArray)
}
