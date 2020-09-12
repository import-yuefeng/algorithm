package main

type optimal struct {
	bestx []int
	x     []int
	bestw int
	cw    int
	w     []int
	n     int
	r     int
}

func main() {

}

func (p *optimal) Init(w []int) {
	p.w = w
	for _, v := range p.w {
		p.r += v
	}
	p.x = make([]int, len(w))
	p.bestx = make([]int, len(w))
	p.cw = 0
	p.bestw = 0
	// p.n =
}

func (p *optimal) findBest(i int) {
	if i > n {

	}
}
