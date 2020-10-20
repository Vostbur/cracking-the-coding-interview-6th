package chapter_3

import "fmt"

type Plate struct {
	index int
	above *Plate
	below *Plate
}

type Pile struct {
	size     int
	capacity int
	top      *Plate
	bottom   *Plate
}

type RowOfPiles struct {
	capacity int
	piles    []*Pile
}

func GetPlate(index int) *Plate {
	return &Plate{index: index}
}

func GetPile(capacity int) *Pile {
	return &Pile{capacity: capacity}
}

func (p *Pile) push(value int) {
	newPlate := GetPlate(value)
	p.size++
	if p.size == 1 {
		p.bottom, p.top = newPlate, newPlate
		return
	}
	p.top.above = newPlate
	newPlate.below = p.top
	p.top = newPlate
}

func (p *Pile) pop() (int, error) {
	if p.top == nil {
		return -1, fmt.Errorf("Cannot pop. Pile is empty.")
	}
	n := p.top
	p.top = p.top.below
	p.size--
	return n.index, nil
}

func (p *Pile) fromBottom() int {
	n := p.bottom
	p.bottom = p.bottom.above
	if p.bottom != nil {
		p.bottom.below = nil
	}
	p.size--
	return n.index
}

func (p Pile) isFull() bool {
	return p.size == p.capacity
}

func GetRowOfPiles(capacity int) *RowOfPiles {
	return &RowOfPiles{capacity: capacity}
}

func (r *RowOfPiles) push(value int) {
	p := r.getLast()
	if p != nil && !p.isFull() {
		p.push(value)
	} else {
		newPile := GetPile(r.capacity)
		newPile.push(value)
		r.piles = append(r.piles, newPile)
	}
}

func (r *RowOfPiles) pop() (int, error) {
	p := r.getLast()
	if p == nil {
		return -1, fmt.Errorf("Cannot pop. There aren't piles.")
	}
	ret, err := p.pop()
	if err != nil {
		return -1, err
	}
	if p.size == 0 {
		r.piles = r.piles[:len(r.piles)-1]
	}
	return ret, nil
}

func (r *RowOfPiles) popAt(pileIndex int) int {
	return r.shift(pileIndex, true)
}

func (r *RowOfPiles) shift(pileIndex int, fromTop bool) int {
	var ret int

	p := r.piles[pileIndex-1]
	if fromTop {
		ret, _ = p.pop()
	} else {
		ret = p.fromBottom()
	}

	if p.size == 0 {
		r.piles = r.piles[:len(r.piles)-1]
	} else if len(r.piles) > pileIndex {
		tmp := r.shift(pileIndex+1, false)
		p.push(tmp)
	}
	return ret
}

func (r RowOfPiles) getLast() *Pile {
	if len(r.piles) == 0 {
		return nil
	}
	return r.piles[len(r.piles)-1]
}
