package main

import "fmt"

func pairWithTargetSum(intput []int, targetSum int) {
	output := []int{-1, -1}
	pairSum := 0
	startPointer := 0
	endPointer := len(intput) - 1

	fmt.Println("\nInput array : ", intput, "Target Sum :", targetSum)

	for startPointer <= endPointer {
		pairSum = intput[startPointer] + intput[endPointer]

		// if current pair sum is greater decrease the end pointer
		if pairSum > targetSum {
			endPointer -=1
		} else if pairSum < targetSum {
			startPointer += 1 // if current pair sum is less increase the start pointer
		} else {
			// target sum in reached, update output
			output[0] = startPointer
			output[1] = endPointer
			// make endpointer negative to break loop
			endPointer = -1
		}
	}

	fmt.Println(output)
}

func main() {
	pairWithTargetSum([]int{1, 2, 3, 4, 6}, 6)
	pairWithTargetSum([]int{2, 5, 9, 11}, 11)
	pairWithTargetSum([]int{2, 5, 9, 11}, 30)	
}