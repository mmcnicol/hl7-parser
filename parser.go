package main

import (
	"fmt"
	//"strings"
)

type SubSubComposite struct {
	value string
}

type SubComposite struct {
	value string
	subSubComposite []SubSubComposite
}

type Composite struct {
	value string
	subComposite []SubComposite
}

type Segment struct {
	composite []Composite
}

func parse(tokens []Token) ([]Segment, error) {
	//fmt.Println("hello!")
	
	fmt.Println("preparing new segment")
	segments := make([]Segment, 0)
	segment := Segment{}
	segment.composite = make([]Composite, 0)
	//var line int
	//level := 1
	//var prevValue string
	subComposite := SubComposite{}
	
	fmt.Println(" prepared new composite")
	composite := Composite{}
	composite.subComposite = make([]SubComposite, 0)
	composite.value = ""
	
    //for _, token := range tokens {
	for i := 0; i < len(tokens); i++ {

		//line = tokens[i].pos.line
		
		if tokens[i].level == 1 {
			if tokens[i].value == "newline" {
				fmt.Println(" newline")
				fmt.Println(" appending segment")
				segments = append(segments, segment)
				
				fmt.Println("preparing new segment")
				segment = Segment{}
				segment.composite = make([]Composite, 0)
			} else if tokens[i].value != "|" {
				composite = Composite{}
				composite.subComposite = make([]SubComposite, 0)
				composite.value = tokens[i].value
			} else if tokens[i].value == "|" {
				fmt.Println("  appending composite on segment")
				segment.composite = append(segment.composite, composite)
				
				fmt.Println(" prepared new composite")
				composite = Composite{}
				composite.subComposite = make([]SubComposite, 0)
				composite.value = ""
			}
		} else if tokens[i].level == 2 {
			if tokens[i].value != "|" {
				if tokens[i].value == "^" && tokens[i].pos.line == 1 && tokens[i].pos.column == 5 {
					subComposite = SubComposite{}
					subComposite.value = tokens[i].value
					subComposite.subSubComposite = make([]SubSubComposite, 0)
					fmt.Println("   appending subComposite on composite")
					composite.subComposite = append(composite.subComposite, subComposite)
				} else if tokens[i].value == "~" && tokens[i].pos.line == 1 && tokens[i].pos.column == 6 {
					subComposite = SubComposite{}
					subComposite.value = tokens[i].value
					subComposite.subSubComposite = make([]SubSubComposite, 0)
					fmt.Println("   appending subComposite on composite")
					composite.subComposite = append(composite.subComposite, subComposite)
				} else if tokens[i].value == "\\&" && tokens[i].pos.line == 1 && tokens[i].pos.column == 8 {
					subComposite = SubComposite{}
					subComposite.value = tokens[i].value
					subComposite.subSubComposite = make([]SubSubComposite, 0)
					fmt.Println("   appending subComposite on composite")
					composite.subComposite = append(composite.subComposite, subComposite)
				} else if tokens[i].value != "^" && tokens[i].value != "~" && tokens[i].value != "\\&" {
					subComposite = SubComposite{}
					subComposite.value = tokens[i].value
					subComposite.subSubComposite = make([]SubSubComposite, 0)
					fmt.Println("   appending subComposite on composite")
					composite.subComposite = append(composite.subComposite, subComposite)
				}
				
			} else if tokens[i].value == "|" {
				fmt.Println("  appending composite on segment")
				segment.composite = append(segment.composite, composite)
				
				fmt.Println("  prepared new composite")
				composite = Composite{}
				composite.subComposite = make([]SubComposite, 0)
				composite.value = ""
			}
		}
		
    }
	if len(segment.composite) > 0 {
		fmt.Println(" appending segment")
		segments = append(segments, segment)
	}
	
	return segments, nil
}
