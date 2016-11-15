package main

/*输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，
则重建二叉树并返回。
*/
import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}
//new一个节点
func NewTreeNode(val int) *TreeNode {
	var result = new(TreeNode)
	result.val = val
	return result
}
//计算一个数的n次
func pow(num int, count int) int {
	sum := num
	for count > 1 {
		sum = sum * num
		count--
	}
	return sum
}
//打印树
func (tn *TreeNode) printTree() {
	queue := list.New()
	queue.PushBack(tn)
	i := 0
	n := 1
	for queue.Len() > 0 {
		i++
		var ptr (*TreeNode)
		ptr = queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if -1 == ptr.val {
			fmt.Print("#")
		} else {
			fmt.Print(ptr.val)
		}
		if i == pow(2, n)-1 {
			fmt.Println()
			n++
		}
		if n > tn.high() {
			break
		}
		if nil != ptr.left {
			queue.PushBack(ptr.left)
		} else {
			queue.PushBack(NewTreeNode(-1))
		}
		if nil != ptr.right {
			queue.PushBack(ptr.right)
		} else {
			queue.PushBack(NewTreeNode(-1))
		}

	}

}
//返回二叉树的高度
func (tn *TreeNode) high() int {
	var ptr = tn
	if ptr.left == nil && ptr.right == nil {
		return 1
	}
	var left, right int
	if ptr.left != nil {
		left = ptr.left.high() + 1
	}
	if ptr.right != nil {
		right = ptr.right.high() + 1
	}
	if left < right {
		return right
	} else {
		return left
	}
}
//返回一个数在数组中的位置
func returnSerialNumber(number int, arr2 []int) int {
	for index, value := range arr2 {
		if number == value {
			return index
		}
	}
	return -1
}
//由前序和后序遍历得到二叉树
func reConstructBinaryTree(arr1 []int, arr2 []int) *TreeNode {
	var tn = NewTreeNode(arr1[0])
	var index = returnSerialNumber(arr1[0], arr2)
	if index != 0 {
		tn.left = reConstructBinaryTree(arr1[1:index+1], arr2[0:index])
	}
	if index != len(arr2)-1 {
		tn.right = reConstructBinaryTree(arr1[index+1:], arr2[index+1:])
	}
	return tn
}

func main() {
	arr1 := []int{1, 2, 4, 7, 3, 5, 6, 8}
	arr2 := []int{4, 7, 2, 1, 5, 3, 8, 6}
	var root *TreeNode
	root = reConstructBinaryTree(arr1, arr2)
	root.printTree()
}
