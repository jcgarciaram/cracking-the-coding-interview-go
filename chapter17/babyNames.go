//https://play.golang.org/p/G966tooIq5

package main

import (
	"fmt"
)

type BabyNameNode struct {
	name 		string
	num  		int
	marked		bool
	nicknames 	[]*BabyNameNode
}

type BabyNameGraph struct {
	nodes		map[string]*BabyNameNode
}

type BabyName struct {
	name	string
	num	int
}

type SynonymName struct {
	name1	string
	name2	string
} 

func buildBabyNameGraph(names []BabyName) BabyNameGraph {
	graph := BabyNameGraph{}
	graph.nodes = make(map[string]*BabyNameNode)
	for i:=0; i < len(names); i++ {
		babyNameNode := BabyNameNode{names[i].name, names[i].num, false, []*BabyNameNode{}}
		graph.nodes[names[i].name] = &babyNameNode
	}
	
	return graph
		
	
}

func (graph *BabyNameGraph) createLinks(syns []SynonymName) {
	for i:= 0; i < len(syns); i++ {
		if tempNode1, ok := graph.nodes[syns[i].name1]; ok {
			if tempNode2, ok := graph.nodes[syns[i].name2]; ok {
				tempNode1.nicknames = append(tempNode1.nicknames, tempNode2)
			}
		}
	}
}


func depthFirstSearch(babyNode *BabyNameNode) int {
	totalNum := babyNode.num
	babyNode.marked = true
	
	for _, syn := range babyNode.nicknames {
		if !syn.marked {
			totalNum += depthFirstSearch(syn)
		}
	}
	
	return totalNum

}
		

func babyNameConsolidation(names []BabyName, syns []SynonymName) []BabyName {
	var retBabyNames []BabyName
	
	babyNameGraph := buildBabyNameGraph(names)
	babyNameGraph.createLinks(syns)
	
	for _, node := range babyNameGraph.nodes {
		if !node.marked {
			num := depthFirstSearch(node)
			retBabyNames = append(retBabyNames, BabyName{node.name, num})
			
		}
	}
	
	return retBabyNames
	
}


	
	

func main() {
	
	var babyNames []BabyName
	babyNames = append(babyNames, BabyName{"John", 5})
	babyNames = append(babyNames, BabyName{"Jon", 3})
	babyNames = append(babyNames, BabyName{"Johnny", 8})
	babyNames = append(babyNames, BabyName{"Chris", 10})
	babyNames = append(babyNames, BabyName{"Kris", 5})
	babyNames = append(babyNames, BabyName{"Joan", 8})
	
	fmt.Println(babyNames)
	
	var syns []SynonymName
	syns = append(syns, SynonymName{"John", "Jon"})
	syns = append(syns, SynonymName{"Jon", "Johnny"})
	syns = append(syns, SynonymName{"Johnny", "John"})
	syns = append(syns, SynonymName{"Chris", "Kris"})
	syns = append(syns, SynonymName{"Kris", "Chris"})
	syns = append(syns, SynonymName{"Joan", "Jane"})
	
	
	fmt.Println(syns)
	
	retBabyNames := babyNameConsolidation(babyNames, syns)
	
	fmt.Println(retBabyNames)
	
}
