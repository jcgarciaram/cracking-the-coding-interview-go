package main

import (
    "fmt"
    "strconv"
)

type Queen struct {
    x   int
    y   int
}

func (q *Queen) conflict(q2 *Queen) bool {
    if q.x == q2.x {
        return true
    }
    
    if q.y == q2.y {
        return true
    }
    
    // Left down
    i := q.x - 1
    j := q.y + 1
    for i >= 0 && j <= 8 {
        if i == q2.x && j == q2.y {
            return true
        }
        i--
        j++
    }
    
    // Left up
    i = q.x - 1
    j = q.y - 1
    for i >= 0 && j >= 0 {
        if i == q2.x && j == q2.y {
            return true
        }
        i--
        j--
    }
    
    // Right up
    i = q.x + 1
    j = q.y - 1
    for i <= 8 && j >= 0 {
        if i == q2.x && j == q2.y {
            return true
        }
        i++
        j--
    }
    
    // Right down
    i = q.x + 1
    j = q.y + 1
    for i <= 8 && j <= 8 {
        if i == q2.x && j == q2.y {
            return true
        }
        i++
        j++
    }
    
    return false
}

func initializeGame(initialPositions string) []Queen {
    
    queens := make([]Queen, 8)
    
    for i := 0; i < len(queens); i++ {
        
            var q Queen
            q.x = i
            
            pos, _ := strconv.Atoi(string(initialPositions[i]))
            
            q.y = pos
            
            queens[i] = q
    }
    
    return queens
    
}

func main() {
    
    games := []string{
        "32583211", 
        "58647561", 
        "35712864", 
        "38647511", 
        "53176462", 
        "77854568", 
    }
    
    /*
    games := []string{  
        "58647561", 
        "35712864", 
        "38647511", 
    }
    */
    //for i := 1; i < 8; i++ {
        //tempGames := make([]string, 3)

        //tempGames[0] = games[0][:i] + games[1][i:]
        //tempGames[1] = games[0][:i] + games[2][i:]
        //tempGames[2] = games[1][:i] + games[2][i:]

        for _, game := range games {

	    
            queens := initializeGame(game)
            
            numFriends := 0
            for i := 0; i < len(queens) - 1; i++ {
                for j := i + 1; j < len(queens); j++ {
                    q := queens[i]
                    q2 := queens[j]
                    
                    if con := q.conflict(&q2); !con {
                        numFriends++
                    }
                }
            }
	    fmt.Println(game, numFriends)
            if numFriends == 28 {
                
		return
            }
            
        }
    //}
}
