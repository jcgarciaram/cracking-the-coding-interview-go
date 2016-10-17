package main

import (
	"fmt"
	"strconv"
)

type parenClause struct {
	oper	rune
	nums	[]int
}

func (pc *parenClause) solve() int {
	ret := 0
	if pc.oper == '+' {
		for _, n := range pc.nums {
			ret += n
		}
	} else if pc.oper == '-' {
		ret = pc.nums[0]
		for _, n := range pc.nums[1:] {
			ret -= n
		}
	} else if pc.oper == '*' {
		ret = 1
		for _, n := range pc.nums {
			ret = ret * n
		}
	} else if pc.oper == '/' {
		ret = pc.nums[0]
		for _, n := range pc.nums[1:] {
			ret = ret/n
		}
	}
	//fmt.Println("Solve:", ret)
	return ret
}

func strEquationParse(eq string) int {

	if _, err := strconv.Atoi(string(eq[0])); err == nil {
		ret, _ := strconv.Atoi(eq)
		return ret
	}
	
	pc := parenClause{}
	
	inNum := false
	inRecEq := false
	operFound := false
	startNum := 0
	for i, c := range eq[1:] {
		strC := string(c)
		
		if (c == '+' || c == '-' || c == '*' || c == '/') && !inRecEq && !operFound {
			pc.oper = c
			operFound = true
		} else if (c == '-') && !inNum && !inRecEq && operFound {
			inNum = true
			startNum = i
		} else if _, err := strconv.Atoi(strC); err == nil && !inNum && !inRecEq {
			inNum = true
			startNum = i
		} else if c == ' ' && inNum == true && !inRecEq {
			inNum = false
			//fmt.Println(startNum, eq[startNum+1:i+1])
			tmpInt, _ := strconv.Atoi(eq[startNum+1:i+1])
			pc.nums = append(pc.nums, tmpInt)
		} else if c == '(' && !inRecEq {
			pc.nums = append(pc.nums, strEquationParse(eq[i+2:]))
			inRecEq = true
		} else if c == ')' && inRecEq {
			inRecEq = true
		} else if c == ')' && !inRecEq {
			break
		}
			
	}
	// fmt.Println(pc)
	return pc.solve()
	
}


func main() {
	str := "75"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( * 1 2 )"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( * 4 3 )"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( * 4 ( + 8 3 9 ) )"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( - 9 ( * 8 2 ( + 6 4 2 ) ) )"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( * -10 ( - 8 5 ( + 6 4 ) ) )"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( * -10 ( - -8 -5 ( + 6 -4 ) ) )"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( * 6 -8 )"
	fmt.Println(str, "=", strEquationParse(str))
}
