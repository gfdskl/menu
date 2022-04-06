package main

type LinkTableNode struct {
	pNext *LinkTableNode
}

type LinkTable struct {
	pHead *LinkTableNode
	pTail *LinkTableNode
	number int
}

func  CreateLinkTable() *LinkTable{
	linkTable := new(LinkTable)
	return linkTable
}

// func DeleteLinkTable(pLinkTble *LinkTable) int {
// 	for pLinkTable != nil {

// 	}
// }


func AddLinkTableNode(linkTable *LinkTable,linkTableNode *LinkTableNode) int{
	linkTableNode.pNext = nil
	// 如果链表的头部为空，尾部也一定为空
	if linkTable.pHead == nil {
		linkTable.pHead = linkTableNode
		linkTable.pTail = linkTableNode
	} else {
		linkTable.pTail.pNext = linkTableNode
		linkTable.pTail = linkTableNode
	}
	linkTable.number += 1
	return 1
}

// 删除节点
func DeleteLinkTableNode(linkTable *LinkTable,linkTableNode *LinkTableNode) int{
	head := linkTable.pHead
	tail := linkTable.pTail
	if linkTableNode == head {
		linkTable.pHead = head.pNext
		linkTable.number -= 1
		if linkTable.number == 0 {
			linkTable.pTail = nil
		}
	} else  {
		preNode := head
		tmpNode := head.pNext
		for tmpNode != nil && tmpNode != linkTableNode {
			preNode = preNode.pNext
			tmpNode = tmpNode.pNext
		}
		if tmpNode != nil {
			preNode.pNext = tmpNode.pNext
			linkTable.number -= 1
			if tmpNode == tail {
				linkTable.pTail = preNode
			}
		}
	}
	return 1
}

func GetLinkTableHead(linkTable *LinkTable) *LinkTableNode{
	return linkTable.pHead
}

// 得到下一个链表节点
func GetNextLinkTableNode(linkTable *LinkTable,linkTableNode *LinkTableNode) *LinkTableNode{
	return linkTableNode.pNext
}

