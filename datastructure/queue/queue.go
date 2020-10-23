/**
* @DateTime   : 2020/10/23 9:37
* @Author     : xumamba
* @Description:
**/
package queue

type Queue []int

func (q *Queue)Enqueue(elem int)  {
	*q = append(*q, elem)
}

func (q *Queue) Dequeue() int {
	elem := (*q)[0]
	*q = (*q)[1:]
	return elem
}

func (q *Queue) Len() int {
	return len(*q)
}

func NewQueue() Queue {
	return make([]int, 0)
}