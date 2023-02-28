package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {

	//want := make([]Token, 0)
	//want = append(want, Token{Position{1,1}, "MSH"})
	
	want := []Segment{
        {
			composite: []Composite{
				{ value: "MSH", subComposite: []SubComposite{} },
				{ 
					value: "", subComposite: []SubComposite{ 
						{ value:"^", subSubComposite: []SubSubComposite{} },
						{ value:"~", subSubComposite: []SubSubComposite{} },
						{ value:"\\&", subSubComposite: []SubSubComposite{} },
					}, 
				},
				{ value: "EPIC", subComposite: []SubComposite{} },
				{ value: "EPICADT", subComposite: []SubComposite{} },
				{ value: "SMS", subComposite: []SubComposite{} },
				{ value: "SMSADT", subComposite: []SubComposite{} },
				{ value: "199912271408", subComposite: []SubComposite{} },
				{ value: "CHARRIS", subComposite: []SubComposite{} },
				{
					value: "", subComposite: []SubComposite{ 
						{ value:"ADT", subSubComposite: []SubSubComposite{} },
						{ value:"A04", subSubComposite: []SubSubComposite{} },						
					}, 
				},
			},
		},
		{
			composite: []Composite{
				{ value: "PID", subComposite: []SubComposite{} },
			},
		},
    }
		
	input := []Token{
        {pos: Position{line:1,column:3}, value:"MSH", level: 1},
        {pos: Position{line:1,column:4}, value:"|", level: 1},
        {pos: Position{line:1,column:5}, value:"^", level: 2},
        {pos: Position{line:1,column:6}, value:"~", level: 2},
		{pos: Position{line:1,column:8}, value:"\\&", level: 2},
		{pos: Position{line:1,column:9}, value:"|", level: 1},
		{pos: Position{line:1,column:13}, value:"EPIC", level: 1},
		{pos: Position{line:1,column:14}, value:"|", level: 1},
		{pos: Position{line:1,column:21}, value:"EPICADT", level: 1},
		{pos: Position{line:1,column:22}, value:"|", level: 1},
		{pos: Position{line:1,column:25}, value:"SMS", level: 1},
		{pos: Position{line:1,column:26}, value:"|", level: 1},
		{pos: Position{line:1,column:32}, value:"SMSADT", level: 1},
		{pos: Position{line:1,column:33}, value:"|", level: 1},
		{pos: Position{line:1,column:45}, value:"199912271408", level: 1},
		{pos: Position{line:1,column:46}, value:"|", level: 1},
		{pos: Position{line:1,column:53}, value:"CHARRIS", level: 1},
		{pos: Position{line:1,column:54}, value:"|", level: 1},
		{pos: Position{line:1,column:57}, value:"ADT", level: 2},
		{pos: Position{line:1,column:58}, value:"^", level: 2},
		{pos: Position{line:1,column:61}, value:"A04", level: 2},
		{pos: Position{line:1,column:62}, value:"|", level: 1},
		{pos: Position{line:1,column:63}, value:"newline", level: 1},
		{pos: Position{line:2,column:3}, value:"PID", level: 1},
		{pos: Position{line:2,column:4}, value:"|", level: 1},
    }
	
	got, err := parse(input)
	
	if err != nil {
		t.Errorf("parse() returned %v ", err.Error())
	} else if !reflect.DeepEqual(want, got) {
		t.Errorf("parse() \nwant %+v \ngot  %+v ", want, got)
	}
	
}
