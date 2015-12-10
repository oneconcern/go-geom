package geom

type Polygon struct {
	geom2
}

var _ T = &Polygon{}

func NewPolygon(layout Layout) *Polygon {
	return NewPolygonFlat(layout, nil, nil)
}

func NewPolygonFlat(layout Layout, flatCoords []float64, ends []int) *Polygon {
	p := new(Polygon)
	p.layout = layout
	p.stride = layout.Stride()
	p.flatCoords = flatCoords
	p.ends = ends
	return p
}

func (p *Polygon) Area() float64 {
	return area2(p.flatCoords, 0, p.ends, p.stride)
}

func (p *Polygon) Clone() *Polygon {
	flatCoords := make([]float64, len(p.flatCoords))
	copy(flatCoords, p.flatCoords)
	ends := make([]int, len(p.ends))
	copy(ends, p.ends)
	return NewPolygonFlat(p.layout, flatCoords, ends)
}

func (p *Polygon) Length() float64 {
	return length2(p.flatCoords, 0, p.ends, p.stride)
}

func (p *Polygon) LinearRing(i int) *LinearRing {
	offset := 0
	if i > 0 {
		offset = p.ends[i-1]
	}
	return NewLinearRingFlat(p.layout, p.flatCoords[offset:p.ends[i]])
}

func (p *Polygon) MustSetCoords(coords [][]Coord) *Polygon {
	if err := p.setCoords(coords); err != nil {
		panic(err)
	}
	return p
}

func (p *Polygon) NumLinearRings() int {
	return len(p.ends)
}

func (p *Polygon) Push(lr *LinearRing) error {
	if lr.layout != p.layout {
		return ErrLayoutMismatch{Got: lr.layout, Want: p.layout}
	}
	p.flatCoords = append(p.flatCoords, lr.flatCoords...)
	p.ends = append(p.ends, len(p.flatCoords))
	return nil
}

func (p *Polygon) SetCoords(coords [][]Coord) (*Polygon, error) {
	if err := p.setCoords(coords); err != nil {
		return nil, err
	}
	return p, nil
}
