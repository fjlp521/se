package linktable

import (
	"strconv"
	"strings"
	"sync"
)

type LinkTableNode struct {
	Val   int
	PNext *LinkTableNode
}

type LinkTable struct {
	pHead, pTail *LinkTableNode
	sumOfNode    int
	mutex        *sync.Mutex
}

func CreateLinkTable() *LinkTable {
	return &LinkTable{
		pHead:     nil,
		pTail:     nil,
		sumOfNode: 0,
		mutex:     &sync.Mutex{},
	}
}

func (lt *LinkTable) DeleteLinkTable() bool {
	if lt == nil {
		return false
	}
	lt.pHead, lt.pTail = nil, nil
	lt.sumOfNode = 0
	return true
}

func (lt *LinkTable) AddLinkTableNode(pNode *LinkTableNode) bool {
	if lt == nil || pNode == nil {
		return false
	}
	pNode.PNext = nil
	lt.mutex.Lock()
	if lt.pHead == nil {
		lt.pHead, lt.pTail = pNode, pNode
	} else {
		lt.pTail.PNext = pNode
		lt.pTail = pNode
	}
	lt.sumOfNode++
	lt.mutex.Unlock()
	return true
}

func (lt *LinkTable) DelLinkTableNode(pNode *LinkTableNode) bool {
	if lt == nil || pNode == nil {
		return false
	}
	lt.mutex.Lock()
	if lt.pHead == pNode {
		lt.pHead = lt.pHead.PNext
		lt.sumOfNode--
		if lt.sumOfNode == 0 {
			lt.pTail = nil
		}
		lt.mutex.Unlock()
		return true
	}
	pTempNode := lt.pHead
	for pTempNode != nil {
		if pTempNode.PNext == pNode {
			if lt.pTail == pNode {
				lt.pTail = pTempNode
			}
			pTempNode.PNext = pTempNode.PNext.PNext
			lt.sumOfNode--
			lt.mutex.Unlock()
			return true
		}
		pTempNode = pTempNode.PNext
	}
	lt.mutex.Unlock()
	return false
}

func (lt *LinkTable) SearchLinkTableNode(Condition func(*LinkTableNode) bool) *LinkTableNode {
	if lt == nil || Condition == nil {
		return nil
	}
	pNode := lt.pHead
	for pNode != nil {
		if Condition(pNode) {
			return pNode
		}
		pNode = pNode.PNext
	}
	return nil
}

func (lt *LinkTable) GetLinkTableHead() *LinkTableNode {
	if lt == nil {
		return nil
	}
	return lt.pHead
}

func (lt *LinkTable) GetNextLinkTableNode(pNode *LinkTableNode) *LinkTableNode {
	if lt == nil || pNode == nil {
		return nil
	}
	pTempNode := lt.pHead
	for pTempNode != nil {
		if pTempNode == pNode {
			return pTempNode.PNext
		}
		pTempNode = pTempNode.PNext
	}
	return nil
}

func (lt *LinkTable) String() string {
	if lt == nil || lt.pHead == nil {
		return "linkTable is null"
	}
	pNode := lt.pHead
	res := &strings.Builder{}
	for pNode != nil {
		res.WriteString(strconv.Itoa(pNode.Val))
		res.WriteString(" ")
		pNode = pNode.PNext
	}
	return res.String()
}
