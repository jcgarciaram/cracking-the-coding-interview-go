package main

import (
	"fmt"
	"strconv"
)

type parenClause struct {
	oper	byte
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
	fmt.Println("Solve:", string(pc.oper), pc.nums, "=", ret)
	return ret
}

func strEquationParseRecursive(eq string) (int, int) {
	
	indexRet := 0
	
	if _, err := strconv.Atoi(string(eq[0])); err == nil {
		ret, _ := strconv.Atoi(eq)
		return ret, len(eq)
	}
	
	pc := parenClause{}
	
	inNum := false
	operFound := false
	startNum := 0
	for i:=1; i < len(eq); i++ {
		c := eq[i]
		strC := string(c)
		fmt.Printf("'%s'\n",string(c))
		
		if (c == '+' || c == '-' || c == '*' || c == '/') && !operFound {
			pc.oper = c
			operFound = true
		} else if (c == '-') && !inNum && operFound {
			inNum = true
			startNum = i
		} else if _, err := strconv.Atoi(strC); err == nil && !inNum {
			inNum = true
			startNum = i
		} else if c == ' ' && inNum == true {
			inNum = false
			//fmt.Println(startNum, eq[startNum:i])
			tmpInt, _ := strconv.Atoi(eq[startNum:i])
			pc.nums = append(pc.nums, tmpInt)
		} else if c == '(' {
			innerParentInt, lenInnerParen := strEquationParseRecursive(eq[i:])
			i += lenInnerParen
			pc.nums = append(pc.nums, innerParentInt)
		} else if c == ')' {
			indexRet = i 
			break
		}
			
	}
	//fmt.Println(pc)
	return pc.solve(), indexRet
	
}

func strEquationParse(eq string) int {
	ret, _ := strEquationParseRecursive(eq)
	return ret
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
	str = "( + ( - 5 3 ) 8 )"
	fmt.Println(str, "=", strEquationParse(str))
	str = "( + ( * 5 ( + 4 5 ) 7 ) 8 )"
	fmt.Println(str, "=", strEquationParse(str))
}
