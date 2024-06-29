// 2192 - All Ancestors of a Node in a Directed Acyclic Graph

package problems

import "container/list"

func bfs(result [][]int, graph map[int][]int, parent int) {
	seen := map[int]bool{
		parent: true,
	}
	queue := list.New()
	queue.PushBack(parent)

	for queue.Len() > 0 {
		current := queue.Front().Value.(int)
		queue.Remove(queue.Front())

		for _, children := range graph[current] {
			if !seen[children] {
				seen[children] = true
				queue.PushBack(children)
				result[children] = append(result[children], parent)
			}
		}
	}
}

func GetAncestors(n int, edges [][]int) [][]int {
	graph := make(map[int][]int)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	result := make([][]int, n)

	for i := 0; i < n; i++ {
		bfs(result, graph, i)
	}

	return result
}
