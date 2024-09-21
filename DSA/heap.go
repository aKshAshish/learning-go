package main

type Heap struct {
	store []int
}

func (h *Heap) add(val int) {
	h.store = append(h.store, val)
	h.heapify()
}

func (h *Heap) heapify() {

}

func main() {
	values := []int{2, 43, 13, 27, 18, 10, 47, 51}
	h := &Heap{store: make([]int, 0)}

	for _, val := range values {
		h.add(val)
	}
}
