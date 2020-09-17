package stack_queue

import (
	"fmt"
	"strconv"
)

/*
栈的特点是后入先出 常用于 DFS 深度搜索
队列一般常用于 BFS 广度搜索，类似一层一层的搜索
*/

/*
min-stack
一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
用两个栈实现，一个最小栈始终保证最小值在顶部
*/
type MinStack struct {
	min, stack []int
}

type Stack struct {
	stack []int
}

func (this *Stack) push(x int) {
	this.stack = append(this.stack, x)
}

func (this *Stack) top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *Stack) pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]

}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) push(x int) {
	min := this.getMin()
	if x < min {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min)
	}
	this.stack = append(this.stack, x)

}

func (this *MinStack) pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) getMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	min := this.min[len(this.min)-1]
	return min
}

/*
波兰表达式
输入: ["2", "1", "+", "3", "*"] > 输出: 9 解释: ((2 + 1) * 3) = 9
通过栈保存原来的元素，遇到表达式弹出运算，再推入结果，重复这个过程
*/
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	var stack Stack
	for _, val := range tokens {
		switch val {
		case "+", "-", "*", "/":
			if len(stack.stack) < 2 {
				return -1
			}
			num1 := (&stack).top()
			(&stack).pop()
			num2 := (&stack).top()
			(&stack).pop()
			switch val {
			case "+":
				(&stack).push(num1 + num2)
			case "-":
				(&stack).push(num1 - num2)
			case "*":
				(&stack).push(num1 * num2)
			case "/":
				(&stack).push(num2 / num1)
			}

		default:
			num, _ := strconv.Atoi(val)
			(&stack).push(num)
		}
	}
	return stack.stack[0]
}

/*
decode string
*/
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			//get loop content
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			//pop '['
			stack = stack[:len(stack)-1]
			var count string
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				count = string(stack[len(stack)-1]) + count
				stack = stack[:len(stack)-1]
			}
			iter, _ := strconv.Atoi(count)
			for iter > 0 {
				insert_temp := temp
				for len(insert_temp) > 0 {
					stack = append(stack, insert_temp[len(insert_temp)-1])
					insert_temp = insert_temp[:len(insert_temp)-1]
				}
				iter--
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)

}

/*
number of island
给定一个由  '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。
一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。
你可以假设网格的四个边均被水包围。
通过深度搜索遍历可能性（注意标记已访问元素
*/
// func numIslands(grid [][]byte) int {
// 	var count int
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func dfs(grid [][]byte, i, j int) int {
// 	if i < 0 || j < 0 || i >= len(grid[0]) || j >= len(grid[0]) {
// 		return 0
// 	}
// 	if grid[i][j] == '1' {
// 		grid[i][j] = 0
// 		return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1) + 1
// 	}
// 	return 0
// }

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
			}
			dfs(grid, i, j)
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = 0
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
}

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
求以当前柱子为高度的面积，即转化为寻找小于当前值的左右两边值
用栈保存小于当前值的左的元素
*/

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	max_area := 0
	for i := 0; i < len(heights); i++ {
		var current_height int
		current_height = heights[i]
		for len(stack) != 0 && current_height < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) != 0 {
				width = i - stack[len(stack)-1] - 1
			}
			max_area = Max(max_area, height*width)
		}
		stack = append(stack, i)
	}
	return max_area

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
01-matrix
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
两个相邻元素间的距离为 1
*/
func updateMatrix(matrix [][]int) [][]int {
	q := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				// 进队列
				point := []int{i, j}
				q = append(q, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for len(q) != 0 {
		// 出队列
		point := q[0]
		q = q[1:]
		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				// 将当前的元素进队列，进行一次BFS
				q = append(q, []int{x, y})
			}
		}
		fmt.Println(matrix)
	}
	return matrix

}
func Test() {
	// min_stack := Constructor()
	// min_stack.push(1)
	// min_stack.push(2)
	// min_stack.push(3)
	// min_stack.pop()
	// fmt.Println(min_stack.stack)
	// fmt.Println(min_stack.getMin())
	// var stack Stack
	// (&stack).push(1)
	// (&stack).push(2)
	// x := (&stack).top()
	// fmt.Println(x)
	// (&stack).pop()
	// fmt.Println(stack.stack)
	var matrix [][]int = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	updateMatrix(matrix)
	// fmt.Println(decodeString(express))

}
