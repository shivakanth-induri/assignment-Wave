package main

import (
	"fmt"
	"golangProject/assignment-Wave/assignment"
)

func main() {
	towers := [][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}
	radius := 2
	fmt.Println(assignment.MaxNetworkQuality(towers, radius))
}
