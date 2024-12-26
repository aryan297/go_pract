package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {

	if len(s.items) == 0 {
		return -1
	}
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove

}
func main() {
	fmt.Println("Queue FIFO")
	Qu()
	fmt.Println("Stack FILO")
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	value1, value2 := two(arr, 10)
	fmt.Println(value1, value2)
	fmt.Println(reverse(arr))
	fmt.Println("change to palindrome")
	fmt.Println(plaindrome("madam"))
	zero := []int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 0, 8, 0, 9, 0}
	moveZero(zero)
	fmt.Println(zero)
	zeroArray := []int{-1, 0, 1, 2, -3, -4}
	fmt.Println(sumZero(zeroArray))
	l := &Tree{}
	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.insert(4)
	l.insert(5)
	l.Display()

}
