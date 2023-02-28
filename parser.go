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
	fmt.Println("hello!")
	
	segments := make([]Segment, 0)
	segment := Segment{}
	segment.composite = make([]Composite, 0)
	line := 1
	level := 1
	var prevValue string
	subComposite := SubComposite{}
	
    //for _, token := range tokens {
	for i := 0; i < len(tokens); i++ {

		line = tokens[i].pos.line
		
		if tokens[0].value=="MSH" && line==1 && tokens[i].pos.column==5 {
			composite := Composite{}
			composite.subComposite = make([]SubComposite, 0)
			composite.value = ""
			subComposite = SubComposite{}
			subComposite.value = "^"
			subComposite.subSubComposite = nil
			composite.subComposite = append(composite.subComposite, subComposite)
			subComposite = SubComposite{}
			subComposite.value = "~"
			subComposite.subSubComposite = nil
			composite.subComposite = append(composite.subComposite, subComposite)
			subComposite = SubComposite{}
			subComposite.value = "\\&"
			subComposite.subSubComposite = nil
			composite.subComposite = append(composite.subComposite, subComposite)
			segment.composite = append(segment.composite, composite)
			i = 6
		}
		
        switch(tokens[i].value) {
			case "|":
				fmt.Println("|, level:", level)
				if level==1 {
					composite := Composite{}
					composite.value = prevValue
					composite.subComposite = nil
					segment.composite = append(segment.composite, composite)					
				} else if level==2 {
				
					composite := Composite{}
					composite.subComposite = make([]SubComposite, 0)
					composite.value = ""
					composite.subComposite = append(composite.subComposite, subComposite)
					subComposite = SubComposite{}
					subComposite.value = prevValue
					subComposite.subSubComposite = nil
					composite.subComposite = append(composite.subComposite, subComposite)
					segment.composite = append(segment.composite, composite)					
				}
				level = 1
			case "^":
				fmt.Println("^")
				if line==1 && tokens[i].pos.column!=5 {
					level = 2
				}
				subComposite = SubComposite{}
				subComposite.value = prevValue
				subComposite.subSubComposite = nil
			case "~":
				fmt.Println("~")
				if line==1 && tokens[i].pos.column!=6 {
					level = 3
				}
			case "\\&":
				fmt.Println("\\&")
				if line==1 && tokens[i].pos.column!=8 {
					level = 4
				}
			case "newline":
				fmt.Println("newline")
				fmt.Println("appending segment")
				segments = append(segments, segment)
				segment = Segment{}
				segment.composite = make([]Composite, 0)
			default:
				fmt.Println("default:", tokens[i].value)
				prevValue = tokens[i].value
		}
    }
	if len(segment.composite) > 0 {
		fmt.Println("appending segment")
		segments = append(segments, segment)
	}
	
	return segments, nil
}
