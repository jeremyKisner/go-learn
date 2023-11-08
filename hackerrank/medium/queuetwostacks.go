package medium

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// In this challenge, you must first implement a queue using two stacks. Then process  queries, where each query is one of the following  types:

// 1 x: Enqueue element  into the end of the queue.
// 2: Dequeue the element at the front of the queue.
// 3: Print the element at the front of the queue.

type Queue struct {
	elements []string
}

func (q *Queue) Enqueue(e string) {
	q.elements = append(q.elements, e)
}

func (q *Queue) Dequeue() {
	if len(q.elements) > 0 {
		q.elements = q.elements[1:]
	}
}

func (q *Queue) Print() {
	if len(q.elements) > 0 {
		fmt.Println(q.elements[0])
	}
}

func QueueTwoStacks() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	queries, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	q := Queue{}
	for i := 0; i < queries; i++ {
		scanner.Scan()
		input := scanner.Text()
		queue := strings.Split(input, " ")
		qType := queue[0]
		if qType == "1" {
			q.Enqueue(queue[1])

		}
		if qType == "2" {
			q.Dequeue()
		}
		if qType == "3" {
			q.Print()
		}
	}
}
