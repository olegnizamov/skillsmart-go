package aisd01_task1

import (
	"testing"
)

func TestCountHasElements(t *testing.T) {

	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)

	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)

	var v int
	v = linkedlist.Count()
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}

func TestCountZero(t *testing.T) {

	var linkedlist = new(LinkedList)
	var v int
	v = linkedlist.Count()
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}

func TestCleanHasElements(t *testing.T) {

	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)

	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)

	linkedlist.Clean()
	if nil != linkedlist.head || nil != linkedlist.tail {
		t.Error("Expected nil LinkedList , got not nil")
	}
}

func TestCleanZero(t *testing.T) {

	var linkedlist = new(LinkedList)
	linkedlist.Clean()
	if nil != linkedlist.head || nil != linkedlist.tail {
		t.Error("Expected nil LinkedList , got not nil")
	}
}

func TestFindZero(t *testing.T) {

	var linkedlist = new(LinkedList)
	var result, err = linkedlist.Find(3)

	if err == nil {
		t.Error("Expected Node{value:-1, next: nil}, got ", result)
	}

}

func TestFindHasElements(t *testing.T) {

	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)

	var result, err = linkedlist.Find(3)
	var answer Node = Node{value: 3, next: &node4}

	if answer.value != result.value && answer.next.value != result.next.value {
		t.Error("Expected Node{value:3, next: &node4}, got ", result)
	}

	if err != nil {
		t.Error("NOT FOUND ??")
	}
}

func TestFindHasElements2(t *testing.T) {

	var node1 Node = Node{value: 0}
	var node2 Node = Node{value: 0}
	var node3 Node = Node{value: 1}
	var node4 Node = Node{value: 1}
	var node5 Node = Node{value: 1}
	var node6 Node = Node{value: 2}
	var node7 Node = Node{value: 2}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node6)
	linkedlist.AddInTail(node7)

	var result Node = linkedlist.Find(1)
	var answer Node = Node{value: 1, next: &node4}
	if answer.value != result.value && answer.next.value != result.next.value {
		t.Error("Expected Node{value:3, next: &node4}, got ", result)
	}
}

func TestSumSameLength(t *testing.T) {

	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)

	var linkedlist2 = new(LinkedList)
	linkedlist2.AddInTail(node1)
	linkedlist2.AddInTail(node2)
	linkedlist2.AddInTail(node3)
	linkedlist2.AddInTail(node4)
	linkedlist2.AddInTail(node5)

	var result = linkedListSum(linkedlist, linkedlist2)
	var answer = new(LinkedList)

	var node1Answer Node = Node{value: 2}
	var node2Answer Node = Node{value: 4}
	var node3Answer Node = Node{value: 6}
	var node4Answer Node = Node{value: 8}
	var node5Answer Node = Node{value: 10}

	answer.AddInTail(node1Answer)
	answer.AddInTail(node2Answer)
	answer.AddInTail(node3Answer)
	answer.AddInTail(node4Answer)
	answer.AddInTail(node5Answer)

	var nodeResult *Node = result.head
	var nodeAnswer *Node = answer.head

	for {

		if nodeResult == nil {
			break
		}

		if nodeResult.value != nodeAnswer.value {
			t.Error("Expected , got ", nodeResult.value)
		}

		nodeResult = nodeResult.next
		nodeAnswer = nodeAnswer.next
	}

}

func TestSumNotSameLength(t *testing.T) {

	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)

	var linkedlist2 = new(LinkedList)
	linkedlist2.AddInTail(node1)
	linkedlist2.AddInTail(node2)
	linkedlist2.AddInTail(node3)
	linkedlist2.AddInTail(node4)
	linkedlist2.AddInTail(node5)
	linkedlist2.AddInTail(Node{value: 6})
	linkedlist2.AddInTail(Node{value: 7})

	var result = linkedListSum(linkedlist, linkedlist2)

	if result.head != nil && result.tail != nil {
		t.Error("Expected nil, got ", result)
	}
}

func TestFindAllHasNot(t *testing.T) {

	var linkedlist = new(LinkedList)
	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)
	var resultNodes []Node = linkedlist.FindAll(100)

	if len(resultNodes) > 0 {
		t.Error("Expected zero count, got ", len(resultNodes))
	}

}

func TestFindAllHasManyElements(t *testing.T) {

	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	var resultNodes []Node = linkedlist.FindAll(5)

	if len(resultNodes) != 3 {
		t.Error("Expected 3 count, got ", len(resultNodes))
	}
}

func TestFindAllHasOneElement(t *testing.T) {

	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	var resultNodes []Node = linkedlist.FindAll(2)

	if len(resultNodes) != 1 {
		t.Error("Expected 1 count, got ", len(resultNodes))
	}
}

func TestInsertFirstNilList(t *testing.T) {

	var node1 Node = Node{value: 1}
	var linkedlist = new(LinkedList)
	linkedlist.InsertFirst(node1)

	if linkedlist.head.value != 1 && linkedlist.tail.value != 1 {
		t.Error("Expected 1, got ", linkedlist.head.value)
	}
}

func TestInsertFirsHasHead(t *testing.T) {

	var nodeNewHead Node = Node{value: 100}
	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	linkedlist.InsertFirst(nodeNewHead)

	if linkedlist.head.value != 100 {
		t.Error("Expected 100, got ", linkedlist.head.value)
	}
}

func TestInsert(t *testing.T) {

	var nodeNewHead Node = Node{value: 100}
	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	var linkedlist = new(LinkedList)
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	linkedlist.AddInTail(node5)
	linkedlist.Insert(&node2, nodeNewHead)

	if linkedlist.head.next.next.value != 100 {
		t.Error("Expected 100, got ", linkedlist.head.value)
	}
}

func TestDeleteHeadNil(t *testing.T) {
	var linkedlist = new(LinkedList)
	linkedlist.Delete(1, false)

	if linkedlist.head != nil {
		t.Error("Expected nil, got ", linkedlist.head.value)
	}
}

func TestDeleteHead(t *testing.T) {
	var linkedlist = new(LinkedList)
	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.Delete(1, false)

	if linkedlist.head.value != 2 {
		t.Error("Expected 2, got ", linkedlist.head.value)
	}
	linkedlist.Delete(2, false)

	if linkedlist.head != nil && linkedlist.tail != nil {
		t.Error("Expected nil, got ", linkedlist.head)
	}
}

func TestDeleteOneTime(t *testing.T) {
	var linkedlist = new(LinkedList)
	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 3}
	var node4 Node = Node{value: 4}
	var node5 Node = Node{value: 5}
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)
	linkedlist.Delete(3, false)

	if linkedlist.head.next.next.value == 3 {
		t.Error("Expected 4, got ", linkedlist.head.value)
	}
}

func TestDeleteManyTime(t *testing.T) {
	var linkedlist = new(LinkedList)
	var node1 Node = Node{value: 1}
	var node2 Node = Node{value: 2}
	var node3 Node = Node{value: 1}
	var node4 Node = Node{value: 1}
	var node5 Node = Node{value: 1}
	linkedlist.AddInTail(node1)
	linkedlist.AddInTail(node2)
	linkedlist.AddInTail(node3)
	linkedlist.AddInTail(node4)
	linkedlist.AddInTail(node5)

	linkedlist.Delete(1, true)

	if linkedlist.head.value != 2 {
		t.Error("Expected 2, got ", linkedlist.head.value)
	}

	if linkedlist.Count() != 1 {
		t.Error("Expected 1 count, got ", linkedlist.Count())
	}
}
