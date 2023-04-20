package greedy

import "fmt"

func LeastInterval(tasks []byte, n int) int {
	return leastInterval(tasks, n)
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 || len(tasks) == 0 {
		return len(tasks)
	}

	cache := make(map[byte]int)

	// assume Task is not sorted
	for _, v := range tasks {
		cache[v]++
	}

	taskHeap := TaskHeap{}
	for k, v := range cache {
		task := Task{id: k, count: v}
		taskHeap = append(taskHeap, &task)
	}

	for i := len(taskHeap) / 2; i >= 0; i-- {
		taskHeap.maxHeapTopDown(i)
	}

	/*
		1.find most Task and solve it first
		2. find 2nd most Task and solve it

		and so on

		until only same Task in the pool
		then execute tasks and add idle time n between 2 tasks
	*/

	//var largestTask, secondLagestTask *Task
	//if len(taskHeap) == 1 {
	//	largestTask = taskHeap[0]
	//	return largestTask.count+n*(largestTask.count-1)
	//}

	res := 0
	topTasks := taskHeap.topKthTask(n + 1)

	for topTasks[0].count != 0 {
		isAllDone := true
		taskCount := 0
		for _, task := range topTasks {
			res++
			if task != nil {
				task.count--
				taskCount++
				if task.count > 0 {
					isAllDone = false
				}
				fmt.Printf("%v->", string(task.id))
			} else {
				fmt.Printf("idle->")
			}

		}

		if isAllDone {
			res += taskCount
		} else {
			res += len(topTasks)
		}

		topTasks = taskHeap.topKthTask(n + 1)
	}

	return res
}

type Task struct {
	id    byte
	count int
}

type TaskHeap []*Task

func (t *TaskHeap) maxHeapTopDown(n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1

	var largest int

	if l < len(*t) && (*t)[l].count > (*t)[n].count {
		largest = l
	} else {
		largest = n
	}

	if r < len(*t) && (*t)[r].count > (*t)[largest].count {
		largest = r
	}

	if largest != n {
		(*t)[largest], (*t)[n] = (*t)[n], (*t)[largest]
		t.maxHeapTopDown(largest)
	}
}

func (t *TaskHeap) topKthTask(k int) []*Task {
	res := make([]*Task, 0)

	if len(*t) < k {
		for _, v := range *t {
			res = append(res, v)
			k--
		}

		for i := 0; i < k; i++ {
			res = append(res, nil)
		}

		return res
	}

	for i := 0; i < k; i++ {
		task := (*t)[0]
		(*t) = (*t)[1:]
		res = append(res, task)
		t.maxHeapTopDown(0)
	}
	return res
}
