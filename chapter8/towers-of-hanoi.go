// Problem 8.5 Towers of Hanoi

// https://play.golang.org/p/D9cqRB2EPB

package main

import (
	"fmt"
)

// Used to be able to define functions for the type []Tower
type TowerSlice []Tower

// Double Linked List
type Tower struct {
	Head *Disk
	NumDisks int
	Tail *Disk
}

// Node in Linked List
type Disk struct {
	Value int
	Next *Disk
	Previous *Disk
}

// Pop a disk from a tower
func (t *Tower) pop() Disk {
	d := t.Head
	t.Head = t.Head.Next
	t.NumDisks -= 1
	return *d
}

// Push a disk unto a tower
func (t *Tower) push(d Disk) {
	if t.NumDisks == 0 {
		t.Head = &d
		t.NumDisks += 1
		t.Tail = &d
	} else {
		d.Next = t.Head
		t.Head.Previous = &d
		t.Head = &d
		t.NumDisks += 1
	}
}

// Print all towers
func (towers TowerSlice) print() {
	for _, t := range towers {
		d := t.Tail
		for i := 0; i < t.NumDisks; i++ {
			fmt.Print(d.Value)
			fmt.Print(" ")
			
			d = d.Previous
			
		}
		fmt.Print("\n")
	}
	fmt.Println("------------------")
}

// Function to solve. 
func towersOfHanoi(towers TowerSlice ) TowerSlice  {
	towers.print()
	
	// If number of disks in first tower is 0, do nothing.
	if towers[0].NumDisks == 0 {
		return towers
	}
	
	// If number of disks in first tower is one, simply move the disk to last tower.
	if towers[0].NumDisks == 1 {
		d := towers[0].pop()
		towers[2].push(d)
		towers.print()
		return towers
	}
	
	// If number of disks is at least two, then call makeMove to move the disks from tower 1 to tower 3 using tower 2 and intermediary.
	towers.makeMove(0, 2, 1, towers[0].NumDisks)
	return towers
}


// Function to move all disks from towerOrigin to towerDest using towerInterm as intermediary.

func (towers TowerSlice) makeMove(towerOrigin, towerDest, towerInterm, numDisks int) {
	if numDisks == 2 {
		d := towers[towerOrigin].pop()
		towers[towerInterm].push(d)
		towers.print()
		d = towers[towerOrigin].pop()
		towers[towerDest].push(d)
		towers.print()
		d = towers[towerInterm].pop()
		towers[towerDest].push(d)
		towers.print()
	} else {
		towers.makeMove(towerOrigin, towerInterm, towerDest, numDisks - 1)
		d := towers[towerOrigin].pop()
		towers[towerDest].push(d)
		towers.print()
		towers.makeMove(towerInterm, towerDest, towerOrigin, numDisks - 1)
	}
}

func main() {
	
	numDisks := 12
	
	
	// Create the towers
	var towers TowerSlice
	for i := 0; i < 3; i++ {
		t := Tower{}
		towers = append(towers, t)
	}
	
	// Push initial disks on first tower
	for i := numDisks; i >= 1; i-- {
		d := Disk{i, nil, nil}
		towers[0].push(d)
	}

	// Solve
	towers = towersOfHanoi(towers)

}
