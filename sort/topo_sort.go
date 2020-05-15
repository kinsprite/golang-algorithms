package sort

type GraphColor uint

const (
	WHITE GraphColor = iota
	GRAY
	BLACK
)

type GraphVertex struct {
	Idx           int
	Name          string
	EdgeVertexIdx []int
	// Sorting data
	Color        GraphColor
	VisitorIdx   int
	FoundTime    int
	FinishedTime int
}

func TopoSort(vs []GraphVertex) []int {
	for i := range vs {
		vs[i].Color = WHITE
		vs[i].VisitorIdx = -1
	}

	time := 0
	pushedCount := 0
	result := make([]int, len(vs))

	for i := range vs {
		if vs[i].Color == WHITE {
			deepFirstVisit(vs, &vs[i], &time, result, &pushedCount)
		}
	}

	return result
}

func deepFirstVisit(vs []GraphVertex, u *GraphVertex, time *int, result []int, pushedCount *int) {
	*time++
	u.FoundTime = *time
	u.Color = GRAY

	for _, idx := range u.EdgeVertexIdx {
		v := &vs[idx]

		if v.Color == WHITE {
			v.VisitorIdx = u.Idx
			deepFirstVisit(vs, v, time, result, pushedCount)
		}
	}

	u.Color = BLACK
	*time++
	u.FinishedTime = *time

	// 最先遍历完成的元素放到队尾
	*pushedCount++
	result[len(result)-(*pushedCount)] = u.Idx
}
