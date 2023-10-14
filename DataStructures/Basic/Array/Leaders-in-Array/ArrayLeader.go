package ArrayLeader

/*
 * Print all the LEADERS in an array.
 * An element is a leader if it is greater than all the elements to its right side. And the rightmost element is always a leader. 
 * Example:
 *	Input: arr[] = {16, 17, 4, 3, 5, 2}, 
 *	Output: 17, 5, 2
 *
 * Approach:
 * 		Find Leader by finding suffix maximum (Scan from right)- 
 *			The idea is to scan all the elements from right to left in an array and keep track of the maximum till now.
 *			When the maximum changes its value, print it.
 *
 *			Complexity:
 *				- Time: O(n)
 *				- Aux. Space: O(1)
 */

import "fmt"

func PrintLeaders(input_arr []int, size int) {
	var out_arr []int

	out_arr = append(out_arr, input_arr[size - 1])
	max_from_right := out_arr[0]

	for i := size - 2; i >= 0; i-- {
		if max_from_right < input_arr[i] {
			out_arr = append(out_arr, input_arr[i])
			max_from_right = input_arr[i]
		}
	 }

	 fmt.Println("Leaders in given array: ", out_arr)
}

func main() {

	input_arr := []int{16, 17, 4, 3, 5, 2}

	fmt.Println("Input array: ", input_arr)

	PrintLeaders(input_arr, len(input_arr))
}