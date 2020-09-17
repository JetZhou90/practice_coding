package btree

/*
定义二叉树
*/
import (
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type BinaryTree struct {
	root  *TreeNode
	count int
}

/*
二叉树遍历
*/
// 前序递归

func preorderTraversal(root *TreeNode, result *[]int) []int {

	if root == nil {
		return *result
	}
	*result = append(*result, root.Val)
	preorderTraversal(root.Left, result)
	preorderTraversal(root.Right, result)

	return *result

}
func merge(x, y []int) []int {
	for _, ele := range y {
		x = append(x, ele)
	}
	return x
}

// 前序非递归
// V3：通过非递归遍历
func preorderTraversal2(root *TreeNode) []int {
	// 非递归
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

//中序递归
func inorderTraversal(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = inorderTraversal(root.Left, result)
	result = append(result, root.Val)
	result = inorderTraversal(root.Right, result)

	return result

}

//中序非递归
// 思路：通过stack 保存已经访问的元素，用于原路返回
func inorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左
		}
		// 弹出
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right
	}
	return result
}

//后序递归
func postorderTraversal(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = postorderTraversal(root.Left, result)
	result = postorderTraversal(root.Right, result)
	result = append(result, root.Val)

	return result

}

//后续非递归
func postorderTraversal2(root *TreeNode) []int {
	// 通过lastVisit标识右子节点是否已经弹出
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

// DFS 深度搜索-从上到下
// V1：深度遍历，结果指针作为参数传入到函数内部
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// DFS 深度搜索-从下向上（分治法）
func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	// 返回条件(null & leaf)
	if root == nil {
		return result
	}
	// 分治(Divide)
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果(Conquer)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func preorderTraversalDFS(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func preorderTraversalDFS2(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

//BFS
func levelOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)

		for i := 0; i < len(queue); i++ {
			level := queue[0]
			queue = queue[1:]
			fmt.Println(level.Val)
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		for _, x := range list {
			result = append(result, x)
		}
		// result = append(result, list)
	}
	return result
}

/*
max depth of tree
给定一个二叉树，找出其最大深度。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

/*
判断是否是高度平衡二叉树
*/
func isBalanced(root *TreeNode) bool {
	if maxDepthBalance(root) == -1 {
		return false
	}
	return true
}
func maxDepthBalance(root *TreeNode) int {
	// check
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// 为什么返回-1呢？（变量具有二义性）
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

/*
非空二叉树最大路径和
分治法，分为三种情况：左子树最大路径和最大，右子树最大路径和最大，左右子树最大加根节点最大，
需要保存两个变量：一个保存子树最大路径和，一个保存左右加根节点和，然后比较这个两个变量选择最大值即
*/
type ResultType struct {
	SinglePath, MaxPath int
}

func maxPathSum(root *TreeNode) int {
	result := walking(root)
	return result.MaxPath
}

func walking(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			0,
			-(1 << 31),
		}
	}
	left := walking(root.Left)
	right := walking(root.Right)
	result := ResultType{}

	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
lowest-common-ancestor-of-a-binary-tree
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
分治法，有左子树的公共祖先或者有右子树的公共祖先，就返回子树的祖先，否则返回根节点
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func Test() {
	var left_lef TreeNode
	left_lef.Val = 2
	var right_lef TreeNode
	right_lef.Val = 4
	var root TreeNode
	root.Val = 3
	root.Left = &left_lef
	root.Right = &right_lef
	result := make([]int, 0)
	// result = preorderTraversalDFS(&root)
	// fmt.Println(result)
	// result = preorderTraversalDFS2(&root)
	// fmt.Println(root)
	preorderTraversal(&root, &result)
	fmt.Println(result)

}
