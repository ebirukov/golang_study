package main

import (
	"fmt"
	"strings"
)

type Node struct {
	i    int
	next *Node
}

func (n *Node) String() string {
	buf := new(strings.Builder)
	cur := n

	for cur != nil {
		if _, err := fmt.Fprintf(buf, "%+v, ", cur.i); err != nil {
			panic(err)
		}
		cur = cur.next
	}

	return buf.String()
}

func drop(node *Node, n int) *Node {
	if node == nil {
		return node
	}

	if node.next == nil {
		return node
	}

	if n <= 0 {
		return node
	}

	return drop(node.next, n-1)
}

func dropNFromEnd(head *Node, n int) *Node {
	head = revertRecursive(nil, head)
	head = drop(head, n)
	return revertRecursive(nil, head)
}

func removeNthFromEnd(head *Node, n int) *Node {
	revHead := revert(head, 0)

	return revert(revHead, n)
}

func revertRecursive(prev *Node, node *Node) *Node {
	if node == nil {
		return node
	}

	if node.next == nil {
		node.next = prev

		return node
	}

	next := revertRecursive(node, node.next)
	node.next = prev

	return next
}

func revert(head *Node, n int) *Node {
	if head == nil {
		return nil
	}
	var prev, next *Node

	cur := head
	i := 1

	for cur.next != nil {
		next = cur.next

		if i != n {
			cur.next = prev
			prev = cur
		}

		cur = next
		i++
	}

	cur.next = prev

	return cur
}
