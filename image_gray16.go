// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package tiff

import (
	"image"
	"image/color"
	"reflect"
)

type imageGray16 struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// newImageGray16 returns a new imageGray16 with the given bounds.
func newImageGray16(r image.Rectangle) *imageGray16 {
	return new(imageGray16).Init(make([]uint8, 2*r.Dx()*r.Dy()), 2*r.Dx(), r)
}

func (p *imageGray16) Init(pix []uint8, stride int, rect image.Rectangle) *imageGray16 {
	*p = imageGray16{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *imageGray16) Pix() []byte           { return p.M.Pix }
func (p *imageGray16) Stride() int           { return p.M.Stride }
func (p *imageGray16) Rect() image.Rectangle { return p.M.Rect }
func (p *imageGray16) Channels() int         { return 1 }
func (p *imageGray16) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *imageGray16) ColorModel() color.Model { return colorGray16Model }

func (p *imageGray16) Bounds() image.Rectangle { return p.M.Rect }

func (p *imageGray16) At(x, y int) color.Color {
	return p.Gray16At(x, y)
}

func (p *imageGray16) Gray16At(x, y int) colorGray16 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorGray16{}
	}
	i := p.PixOffset(x, y)
	return pGray16At(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *imageGray16) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*2
}

func (p *imageGray16) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorGray16Model.Convert(c).(colorGray16)
	pSetGray16(p.M.Pix[i:], c1)
	return
}

func (p *imageGray16) SetGray16(x, y int, c colorGray16) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetGray16(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *imageGray16) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &imageGray16{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(imageGray16).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *imageGray16) Opaque() bool {
	return true
}