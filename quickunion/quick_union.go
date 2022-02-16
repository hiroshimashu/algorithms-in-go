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

func (qu *QuickUnion) Find(p, q int) bool {
	return false
}

func (qu *QuickUnion) Union(p, q int) {

}