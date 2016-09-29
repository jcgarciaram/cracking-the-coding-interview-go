package main

import (
	"fmt"
	"errors"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	start Point
	end Point
}

func getSlopeOffset(l Line) (float64, float64) {
	m := (l.start.y - l.end.y) / (l.start.x - l.end.x)
	b := l.start.y - (l.start.x*m)
	return m, b
} 

func pointInLine(l Line, p Point) bool {
	return p.x >= min(l.start.x, l.end.x) && p.x <= max(l.start.x, l.end.x) && p.y >= min(l.start.y, l.end.y) && p.y <= max(l.start.y, l.end.y)

func intersection(a Line, b Line) (Point, error) {

	var p Point
	
	aM, aB := getSlopeOffset(a)
	bM, bB := getSlopeOffset(b)
	
	if aM == bM {
		if aB == bB {
			return p, errors.New("Lines are the same.")
		}
		return p, errors.New("Lines are parallel. They will never intersect.")
	}
	
	p.x = (aB - bB)/(bM - aM)
	p.y = (aM*p.x) + aB
	
	if pointInLine(a, p) && pointInLine(b, p) {
		return p, nil
	}
	
	return p, errors.New("Intersection point is not part of one of the line segments")
	
	
}

func main() {
	a := Line{
		Point{3, 3},
		Point{1, 1},
	}
	b := Line{
		Point{0, 2},
		Point{1, 1},
	}
	if inters, err := intersection(a, b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inters)
	}
}
