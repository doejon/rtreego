package rtreego

import (
	"image"
	"image/color"
	"math"
)

// ImgRGBA is used to add some drawing funcs to default rgba image
type ImgRGBA struct {
	*image.RGBA
	fact    float64 // factor to multiply each point with
	colNode color.RGBA
	colEdge color.RGBA
}

// Img is the default image interface to draw an image
type Img interface {
	Rect(x0 int, y0 int, x1 int, y1 int, col color.Color)
}

// NewImgRGBA creates a new RGBA image
func NewImgRGBA(r image.Rectangle, fact float64) *ImgRGBA {
	i := &ImgRGBA{
		fact: fact,
		colNode: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0xFF,
			A: 0xFF,
		},
		colEdge: color.RGBA{
			R: 0xFF,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
	}
	i.RGBA = image.NewRGBA(r)
	return i
}

// HLine draws a horizontal line
func (i *ImgRGBA) hline(x1, y, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		i.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func (i *ImgRGBA) vline(x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		i.Set(x, y1, col)
	}
}

// Rect draws a rectangle
func (i *ImgRGBA) Rect(x0 int, y0 int, x1 int, y1 int, col color.Color) {
	// draw two horizontal lines
	i.hline(x0, y0, x1, col)
	i.hline(x0, y1, x1, col)
	i.vline(x0, y0, y1, col)
	i.vline(x1, y0, y1, col)
}

// normalizeX normalizes an x value to fit on image
func (i *ImgRGBA) normalize(val float64) int {
	val *= i.fact // multiply with factor
	return int(math.Ceil(val))
}

// PutRect puts a rectangle to image
func (i *ImgRGBA) PutRect(r *Rect, col color.Color) {
	i.Rect(
		i.normalize(r.p[0]),
		i.normalize(r.p[1]),
		i.normalize(r.q[0]),
		i.normalize(r.q[1]),
		col)
}

// PutNode puts a node to image
func (i *ImgRGBA) PutNode(n *node) {
	for _, v := range n.entries {
		i.PutEntry(v)
	}
}

// PutEntry puts an entry
func (i *ImgRGBA) PutEntry(e *entry) {
	if e.child != nil {
		i.PutRect(e.bb, i.colNode)
		i.PutNode(e.child)
	} else {
		i.PutRect(e.bb, i.colEdge)
	}
}
