package quickunion

type QuickUnion struct {
	ids[] int
}

func New(n int) *QuickUnion {
	// Create QU 
	qu := &QuickUnion{}
	qu.ids = make([]int, n)
	// Initialize ids
	for i := 0; i < n; i++ {
		qu.ids[i] = i
	}

	return qu
}

func (qu *QuickUnion) Root(i int) int {
	for (i != qu.ids[i]) {
		i = qu.ids[i]
	}
	return i 
}

func (qu *QuickUnion)Connected(p, q int) bool {
	return qu.Root(p) == qu.Root(q)
}

func (qu *QuickUnion) Union(p, q int) {
	i := qu.Root(p)
	j := qu.Root(q)
	qu.ids[i] = j
}