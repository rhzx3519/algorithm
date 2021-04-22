package minimal_spanning_tree

import (
	"github.com/rhzx3519/algorithm/graph"
	"math"
)

type PrimMST struct {
	adj [][]graph.Edge	// 邻接表
}

// 生成最小生成树
func (p *PrimMST) get() []graph.Edge {
	var mst = []graph.Edge{}
	var N = len(p.adj)
	visited := make([]bool, N)
	dis := make([]int, N)	// 最小生成树到剩余节点的距离
	for i := 0; i < N; i++ {
		dis[i] = math.MaxInt32
	}

	for i := 0; i < N; i++ {
		var minV = -1
		for j := 0; j < N; j++ {
			if visited[j] {
				continue
			}
			if minV == -1 || dis[j] < dis[minV] {
				minV = j
			}
		}
		visited[minV] = true
		for _, e := range p.adj[minV] { // 更新最小生成树到剩余节点的距离
			if !visited[e.V] && dis[e.V] > e.Cost {
				dis[e.V] = e.Cost
			}
		}
	}

	return mst
}

