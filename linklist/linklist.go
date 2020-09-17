package linklist

import (
	"fmt"
)

/*
定义单链表
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val          int
	Next, Random *Node
}

func Test() {
	var head ListNode
	head.Val = 9
	var node1 ListNode
	node1.Val = 3
	var node2 ListNode
	node2.Val = 3
	var node3 ListNode
	node3.Val = 9
	var node4 ListNode
	node4.Val = 9
	// var node5 ListNode
	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	// node3.Next = &node4
	// node4.Next = &node1
	fmt.Println(printHead(&head))
	fmt.Println(isPalindrome(&head))
	// prev := sortList(&head)
	// fmt.Println(isPalindrome(&head))

}

/*
 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。 要求返回这个链表的 深拷贝。
1、hash 表存储指针，2、复制节点跟在原节点后面
复制节点为新的节点指针，并非单纯复制源节点
*/
func DeepcopyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	cur := head
	//将clone节点放置原节点后面，clone节点为新节点指针，而非源节点
	for cur != nil {
		clone := &Node{Val: cur.Val, Next: cur.Next}
		temp := cur.Next
		cur.Next = clone
		cur = temp
	}
	cur = head
	// 指针处理
	for cur != nil {
		if cur.Random != nil {
			// 即clone节点的random指向源节点random指向的节点的clone节点
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next //跳过clone节点
	}
	cur = head
	cloneHead := cur.Next // 去掉第一位的源节点
	for cur != nil && cur.Next != nil {
		// 去掉clone源节点
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	return cloneHead
}

/*
Cycle detection
给定一个链表，判断链表中是否有环。
快慢指针，快慢指针相同则有环，
证明：如果有环每走一步快慢指针距离会减 1
*/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

/*
给定一个链表，返回链表开始入环的第一个节点。  如果链表无环，则返回  null。
快慢指针，快慢相遇之后，慢指针回到头，快慢指针步调一致一起移动，相遇点即为入环点
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

/*
palindrome-linked-list
判断一个链表是否为回文链表。
*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	// fast如果初始化为head.Next则中点在slow.Next
	// fast初始化为head,则中点在slow
	fast := head.Next
	for fast != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tail := reverseList(slow.Next)
	slow.Next = nil
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true

}

/*
merge-two-sorted-lists
将两个升序链表合并为一个新的升序链表并返回。
新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: 0}
	head := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	// 连接l1 未处理完节点
	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	// 连接l2 未处理完节点
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummyNode.Next
}

/*
partition-list
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于  x  的节点都在大于或等于  x  的节点之前。
将大于 x 的节点，放到另外一个链表，最后连接这两个链表
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	headDummy.Next = head
	head = headDummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 移除<x节点
			t := head.Next
			head.Next = head.Next.Next
			// 放到另外一个链表
			tail.Next = t
			tail = tail.Next
		}
	}
	// 拼接两个链表
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}

/*
sort-list
归并排序，找中点和合并操作
*/
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}
func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	// 快指针先为nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil
	left := mergeSort(head)
	right := mergeSort(tail)
	result := mergeTwoList(left, right)
	return result
}

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
*/
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		// 全部删除完再移动到下一个元素
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中,没有重复出现的数字。
用 dummy node 辅助删除
*/
func deleteDuplicatesWithoutNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyNode := &ListNode{Val: 0, Next: head}
	head = dummyNode
	var rmNodeval int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmNodeval = head.Next.Val
			for head.Next != nil && head.Next.Val == rmNodeval {
				head.Next = head.Next.Next
			}

		} else {
			head = head.Next
		}
	}
	return dummyNode.Next
}

/*
反转一个单链表
用一个 prev 节点保存向前指针，temp 保存向后的临时指针
*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

/*
反转从位置  m  到  n  的链表。请使用一趟扫描完成反转。
*/
func reverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil {
		return head
	}
	dummyNode := &ListNode{Val: 0, Next: head}
	head = dummyNode
	var i int = 0
	var pre *ListNode
	for i < m {
		pre = head
		head = head.Next
		i++
	}
	var j int = i
	var next *ListNode
	var mid = head
	for head != nil && j <= n {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		j++
	}
	pre.Next = next
	mid.Next = head
	return dummyNode.Next
}

func reverseBetween2(head *ListNode, m, n int) *ListNode {
	if head == nil {
		return head
	}
	count := int(n-m) + 1
	dummyNode := &ListNode{Val: 0, Next: head}
	head = dummyNode
	var pre *ListNode
	var next *ListNode
	var mid *ListNode
	for head != nil && count > 0 {
		if m > 0 {
			pre = head
			head = head.Next
			m--
			mid = head
		} else {
			count--
			temp := head.Next
			head.Next = next
			next = head
			head = temp
		}
	}
	pre.Next = next
	mid.Next = head
	return dummyNode.Next
}

func printHead(head *ListNode) []int {
	result := make([]int, 0)
	if head == nil {
		return result
	}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
