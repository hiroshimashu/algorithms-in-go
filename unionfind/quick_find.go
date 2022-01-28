package quickfind

type QuickFind struct {
	ids []int
}

func New(n int) *QuickFind {
	qf := &QuickFind{}
	qf.ids = make([]int, n)
	
	for i := 0; i < n; i++ {
		qf.ids[i] = i
	}

	return qf
}

func (q *QuickFind) Connected(a, b int) bool{
	return q.ids[a] == q.ids[b]
}

func (q *QuickFind) Union(a, b int) {
	aid := q.ids[a]
	bid := q.ids[b]

	for i := 0; i < len(q.ids); i++ {
		if (q.ids[i] == aid) {
			q.ids[i] = bid
		}
	}

}