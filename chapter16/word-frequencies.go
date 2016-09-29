//https://play.golang.org/p/Ac4NBoXzYj

package main

import (
	"fmt"
	"errors"
	"strings"
)

func main() {
	book := strings.Split("A book is a set of written, printed, illustrated, or blank sheets, made of ink, paper, parchment, or other materials, fastened together to hinge at one side. A single sheet within a book is a leaf, and each side of a leaf is a page. A set of text-filled or illustrated pages produced in electronic format is known as an electronic book, or e-book. Books may also refer to works of literature, or a main division of such a work. In library and information science, a book is called a monograph, to distinguish it from serial periodicals such as magazines, journals, or newspapers. The body of all written works including books is literature. In novels and sometimes other types of books (for example, biographies), a book may be divided into several large sections, also called books (Book 1, Book 2, Book 3, and so on). An avid reader of books is a bibliophile or colloquially, bookworm.", " ")
	word := "illustrated"
	
	if count, err := wordFrequencies(book, word); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}
}

func wordFrequencies(book []string, word string) (int, error) {
	if word == "" {
		return 0, errors.New("Word is not defined")
	} 
	
	if book == nil {
		return 0, errors.New("Book is not defined")
	}
	
	wordMap := createWordMap(book)
	
	count, ok := wordMap[strings.ToLower(word)]
	if !ok {
		return 0, nil
	}
	return count, nil
	
}


func createWordMap(book []string) map[string]int {
	wordMap := make(map[string]int)
	
	
	
	for _, w := range book {
		w = strings.ToLower(w)
	
		c, ok := wordMap[w]
		if !ok {
			wordMap[w] = 1
		} else {
			wordMap[w] = 1 + c
		}
	}
	return wordMap
}
