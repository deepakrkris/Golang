/**
* The Problem:
*
* We have a list of tasks. Each task can depend on other tasks. 
* For example if task A depends on task B then B should run before A.
* 
* Implement the "getTaskWithDependencies" method such that it returns a list of task names in the correct order.
* 
* For example if I want to execute task "application A", the method should return a list with ["storage", "mongo", "application A"].
*
* List of Tasks:
*
*   - name: application A
*     dependsOn:
*       - mongo
*
*   - name: storage
*
*   - name: mongo
*     dependsOn:
*       - storage
*
*   - name: memcache
*
*   - name: application B
*     dependsOn:
*       - memcache
*
* The Golang program is expected to be executed succesfully.
*/

package main

import "fmt"
import "reflect"

type Queue[T comparable] struct {
	bucket []T
}

func (Q *Queue[T]) dequeue () T {
	val := Q.bucket[len(Q.bucket) - 1]
	Q.bucket = Q.bucket[ : len(Q.bucket) - 1]
	return val
}

func (Q *Queue[T]) enqueue (x T) {
	Q.bucket = append(Q.bucket, x)
}

func (Q *Queue[T]) size() int {
	return len(Q.bucket)	
}

func getTaskWithDependencies(tasks []task, dependsOn string) []string {
  graph := map[string] []string {}
	indegree := map[string] int {}

  // init graph
	for _, t := range tasks {
		graph[t.Name] = []string{}
		indegree[t.Name] = 0
	}

	// build graph
	for _, t := range tasks {
		edges := t.DependsOn
		for _, target := range edges {
			graph[t.Name] = append(graph[t.Name], target)
			indegree[target] += 1
		}
	}
 
  top_order := []string {}

	// init topology Queue
	Q := &Queue[string] {
		bucket : make([]string, 0),
	}

	// process from source node
	Q.enqueue(dependsOn)
	// reset indegree of source node
	indegree[dependsOn] = 0

	//process topology from source (child) to root of dependency tree
  for Q.size() > 0 {
		source := Q.dequeue()
		// prefix new node in dependency tree to topology order
		top_order = append([]string { source } , top_order...)

		for _, target := range graph[source] {
			indegree[target] -= 1
			if indegree[target] == 0 {
				Q.enqueue(target)
			}
		}
	}

	// all nodes from source to root processed , return order
  return top_order
}

func main() {
  verify(
    []string{"storage", "mongo", "application A"},
    getTaskWithDependencies(getTasks(), "application A"),
  )
}

type task struct {
  Name      string
  DependsOn []string
}

func newTask(name string, dependsOn []string) task {
  return task{
    Name: name,
    DependsOn: dependsOn,
  }
}

func getTasks() []task {
  return []task{
    newTask("application A", []string{"mongo"}),
    newTask("storage", []string{}),
    newTask("mongo", []string{"storage"}),
    newTask("memcache", []string{}),
    newTask("application B", []string{"memcache"}),
  }
}

// ===== helper methods =====

func verify(expected, actual []string) {
  if !reflect.DeepEqual(expected, actual) {
    fmt.Printf("❌ Failed test!\n")
    fmt.Printf("  expected: %s\n", expected)
    fmt.Printf("  actual: %s\n", actual)
    return
  }
  
  fmt.Println("✅ Passed test!")
}
