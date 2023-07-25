package main

import (
	"fmt"
)

type Queue[T comparable] struct {
	bucket []T
}

func (q *Queue[T]) enqueue (x T) {
	q.bucket = append(q.bucket, x) 
}

func (q *Queue[T]) dequeue () T {
	val := q.bucket[len(q.bucket) - 1]
	q.bucket = q.bucket[ : len(q.bucket) - 1 ]
	return val
}

func (q *Queue[T]) size () int {
	return len(q.bucket)
}

func top_sort( dependencies [][]rune ) []string {
	graph := map[rune][]rune {}
	indegree := map[rune]int {}

	for _, edge := range dependencies {
		target := edge[0]
		source := edge[1]
		if _ , ok := graph[source] ; !ok {
			graph[source] = make([]rune, 0)
			indegree[source] = 0
		}
		if _ , ok := graph[target] ; !ok {
			graph[target] = make([]rune, 0)
			indegree[target] = 0
		}
		//fmt.Println(string(source), string(target))
		graph[source] = append(graph[source], target)
		indegree[target] += 1
	}

	result := []string {}

	Q := &Queue[rune] {
		bucket : make([]rune, 0),
	}

    //fmt.Println(graph, indegree)

	for k, _ := range graph {
		if indegree[k] == 0 {
	        //fmt.Println(string(k))
			Q.enqueue(k)
		}
	}

	//fmt.Println(Q)

	for Q.size() > 0 {
		source := Q.dequeue()
		result = append(result, string(source))

		for _, target := range graph[source] {
			indegree[target] -= 1
			if indegree[target] == 0 {
				Q.enqueue(target)
			}
		}
	}

	return result
}

func main() {
	dependencies := [][][]rune{
		{
			{'B', 'A'},
			{'C', 'A'},
			{'D', 'C'},
			{'E', 'D'},
			{'E', 'B'},
		},
		{
			{'B', 'A'},
			{'C', 'A'},
			{'D', 'B'},
			{'E', 'B'},
			{'E', 'D'},
			{'E', 'C'},
			{'F', 'D'},
			{'F', 'E'},
			{'F', 'C'},
		},
		{
			{'A', 'B'},
			{'B', 'A'},
		},
		{
			{'B', 'C'},
			{'C', 'A'},
			{'A', 'F'},
		},
		{
			{'C', 'C'},
		},
	}

	for _, prerequisites := range dependencies {
		result := top_sort(prerequisites)
		fmt.Println(result)
	}
}
