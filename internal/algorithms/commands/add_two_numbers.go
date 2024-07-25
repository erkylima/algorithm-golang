package commands

import (
	"strconv"

	"github.com/erkylima/allGolangRithm/internal/algorithms/entity"
)

type OnAddTwoNumbers struct {
	L1 *entity.ListNode
	L2 *entity.ListNode
}

func (two *OnAddTwoNumbers) Execute() *entity.ListNode {
	resto := 0
	result := &entity.ListNode{}
	l3 := result

	for two.L1 != nil || two.L2 != nil {
		n1 := 0
		if two.L1 != nil {
			n1 = two.L1.Val
			two.L1 = two.L1.Next
		}
		n2 := 0
		if two.L2 != nil {
			n2 = two.L2.Val
			two.L2 = two.L2.Next
		}
		temp := n1 + n2 + resto
		if temp >= 10 {
			resto = temp / 10
			temp = temp - 10
		} else {
			resto = 0
		}

		l3.Next = &entity.ListNode{
			Val: temp,
		}
		l3 = l3.Next
	}
	if resto != 0 {
		l3.Next = &entity.ListNode{
			Val: resto,
		}
	}

	return result.Next
}

func (two *OnAddTwoNumbers) Stringfy() string {
	result := two.Execute()

	var res string
	for result != nil {
		text := strconv.Itoa(result.Val)
		res = text + res
		result = result.Next
	}
	return res
}
