package main


/*把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。*/


import "fmt"

func minNumberInRotateArray(arr []int) int {

	if 0 == len(arr) {
		return 0
	}

	pre := arr[0]
	for _, value := range arr[1:] {
		if pre > value {
			return value
		} else {
			pre = value
		}
	}

	return arr[0]
}

func main(){
	arr :=[]int{3,4,5,1,2}
	fmt.Println(minNumberInRotateArray(arr))
}