/* 
Problem 8.5 Towers of Hanoi

https://play.golang.org/p/D9cqRB2EPB

Usage:
  -n int
        number of disks (default 3)
  -t int
        delay between moves animation (default 500) 
*/

package main

import (
	"fmt"
    "os"
    "strconv"
    "bufio"
    "time"
    "os/exec"
    "flag"
)

var (
    f = bufio.NewWriter(os.Stdout)
    numDisks int
    numMoves = 0
    timeDelay int
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
func (ts *TowerSlice) pop(towerIndex int) Disk {
	d := (*ts)[towerIndex].Head
	(*ts)[towerIndex].Head = (*ts)[towerIndex].Head.Next
	(*ts)[towerIndex].NumDisks -= 1
	return *d
}

// Push a disk unto a tower
func (ts *TowerSlice) push(d Disk, towerIndex int, printBool bool) {
    
	if (*ts)[towerIndex].NumDisks == 0 {
    
		(*ts)[towerIndex].Tail = &d
        
	} else {
		d.Next = (*ts)[towerIndex].Head
		(*ts)[towerIndex].Head.Previous = &d		
	}
    
    
    (*ts)[towerIndex].Head = &d
    (*ts)[towerIndex].NumDisks += 1
    
    
    
    if printBool {
        numMoves += 1
        ts.print(true)
    }
}

// Print all towers
func (ts *TowerSlice) print(replace bool) {
    
    if replace {overwriteLinesTerminal(numDisks+2)}
    
    
    currDisks := []*Disk{(*ts)[0].Head, (*ts)[1].Head, (*ts)[2].Head}
    
    for currLine := numDisks; currLine >= 1;  currLine--{
        
        for i := 0; i < 3; i ++ {
            
            if (*ts)[i].NumDisks >= currLine {
                fmt.Print(currDisks[i].Value)
                currDisks[i] = currDisks[i].Next
            } else {
                fmt.Print(" ")
            }
            fmt.Print("   ")
            
        }
        fmt.Print("        \n")
        
    }
    fmt.Printf("---------------\nTotal Moves: %d\n", numMoves)
    time.Sleep(time.Millisecond*time.Duration(timeDelay))
}

// Function to solve. 
func towersOfHanoi(ts TowerSlice) TowerSlice  {
	ts.print(false)

	
	// If number of disks in first tower is one, simply move the disk to last tower.
	if ts[0].NumDisks == 1 {
		d := ts.pop(0)
		ts.push(d, 2, true)
		
		return ts
	}
	
	// If number of disks is at least two, then call makeMove to move the disks from tower 1 to tower 3 using tower 2 and intermediary.
	ts.makeMove(0, 2, 1, ts[0].NumDisks)
	return ts
}


// Function to move all disks from towerOrigin to towerDest using towerInterm as intermediary.

func (towers *TowerSlice) makeMove(towerOrigin, towerDest, towerInterm, numDisks int) {
	if numDisks == 2 {
		d := towers.pop(towerOrigin)
		towers.push(d, towerInterm, true)
		
		d = towers.pop(towerOrigin)
		towers.push(d, towerDest, true)
		
		d = towers.pop(towerInterm)
		towers.push(d, towerDest, true)
		
	} else {
		towers.makeMove(towerOrigin, towerInterm, towerDest, numDisks - 1)
        
		d := towers.pop(towerOrigin)
		towers.push(d, towerDest, true)
		
		towers.makeMove(towerInterm, towerDest, towerOrigin, numDisks - 1)
	}
}


func overwriteLinesTerminal(numLines int) {
    cmd := exec.Command("echo", "-en", "\\e[" + strconv.Itoa(numLines)+ "A")  
	cmd.Stdout = f    
    if err := cmd.Run(); err != nil {
        fmt.Println("Error echoing: ", err)
    }
    
    f.Flush()
}




func main() {
    
    flag.IntVar(&numDisks, "n", 3, "number of disks")
	flag.IntVar(&timeDelay, "t", 500, "delay between moves animation")
    
    flag.Parse()
    
    
    /* var err error
    if len(os.Args) == 1 {
        fmt.Println("\nHey! Just FYI, you can define the number of disks to start with by passing an argument for the number of disks you like:\n\nExample: towers-of-hanoi 3\n\nYou can also define the difference in time between moves by defining a secon")
        
    } else if len(os.Args) == 2 {
    
        if numDisks, err = strconv.Atoi(os.Args[1]); err != nil || numDisks <= 0 {
            fmt.Println("\nWhoops! The argument passed needs to be a positive number :) ")
            return
        }
        
        
        
        
    } */
    
    fmt.Printf("\nRunning towers of hanoi with %d initial disks\n\n", numDisks)
        
        
        
        
    
	
	
	
	// Create the towers
	var ts TowerSlice
	for i := 0; i < 3; i++ {
		t := Tower{}
		ts = append(ts, t)
	}
	
	// Push initial disks on first tower
	for i := numDisks; i >= 1; i-- {
		d := Disk{i, nil, nil}
		ts.push(d, 0, false)
	}

	// Solve
	ts = towersOfHanoi(ts)
    fmt.Println()
}

