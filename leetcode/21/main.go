package main

func main() {
	// list1 := []int{1, 2, 4}
	// list2 := []int{1, 3, 4}
	// result := mergeTwoLists(list1, list2)
	// fmt.Println(result)
}

// func shift(arr []int) (int, []int) {
// 	if len(arr) == 1 {
// 		return arr[0], []int{}
// 	}

// 	return arr[0], arr[1:]
// }

// func mergeTwoLists(list1 []int, list2 []int) []int {
// 	longerList := list1
// 	otherList := list2
// 	if len(list1) < len(list2) {
// 		longerList = list2
// 		otherList = list1
// 	}

// 	result := []int{}

// 	for len(longerList) > 0 {
// 		var shifted int
// 		var otherShifted int

// 		if len(otherList) == 0 {
// 			result = append(result, longerList...)
// 			break
// 		}

// 		shifted, longerList = shift(longerList)
// 		otherShifted, otherList = shift(otherList)
// 		if shifted >= otherShifted {
// 			result = append(result, otherShifted)
// 			longerList = append([]int{shifted}, longerList...)
// 		} else {
// 			result = append(result, shifted)
// 			otherList = append([]int{otherShifted}, otherList...)
// 		}

// 	}

// 	return result
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
