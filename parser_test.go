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
				{ value: "MSH", subComposite: nil },
				{ 
					value: "", subComposite: []SubComposite{ 
						{ value:"^", subSubComposite:nil },
						{ value:"~", subSubComposite:nil },
						{ value:"\\&", subSubComposite:nil },
					}, 
				},
				{ value: "EPIC", subComposite: nil },
				{ value: "EPICADT", subComposite: nil },
				{ value: "SMS", subComposite: nil },
				{ value: "SMSADT", subComposite: nil },
				{ value: "199912271408", subComposite: nil },
				{ value: "CHARRIS", subComposite: nil },
				{
					value: "", subComposite: []SubComposite{ 
						{ value:"ADT", subSubComposite:nil },
						{ value:"A04", subSubComposite:nil },						
					}, 
				},
			},
		},
		{
			composite: []Composite{
				{ value: "PID", subComposite: nil},
			},
		},
    }
		
	input := []Token{
        {pos: Position{line:1,column:3}, value:"MSH"},
        {pos: Position{line:1,column:4}, value:"|"},
        {pos: Position{line:1,column:5}, value:"^"},
        {pos: Position{line:1,column:6}, value:"~"},
		{pos: Position{line:1,column:8}, value:"\\&"},
		{pos: Position{line:1,column:9}, value:"|"},
		{pos: Position{line:1,column:13}, value:"EPIC"},
		{pos: Position{line:1,column:14}, value:"|"},
		{pos: Position{line:1,column:21}, value:"EPICADT"},
		{pos: Position{line:1,column:22}, value:"|"},
		{pos: Position{line:1,column:25}, value:"SMS"},
		{pos: Position{line:1,column:26}, value:"|"},
		{pos: Position{line:1,column:32}, value:"SMSADT"},
		{pos: Position{line:1,column:33}, value:"|"},
		{pos: Position{line:1,column:45}, value:"199912271408"},
		{pos: Position{line:1,column:46}, value:"|"},
		{pos: Position{line:1,column:53}, value:"CHARRIS"},
		{pos: Position{line:1,column:54}, value:"|"},
		{pos: Position{line:1,column:57}, value:"ADT"},
		{pos: Position{line:1,column:58}, value:"^"},
		{pos: Position{line:1,column:61}, value:"A04"},
		{pos: Position{line:1,column:62}, value:"|"},
		{pos: Position{line:1,column:63}, value:"newline"},
		{pos: Position{line:2,column:3}, value:"PID"},
		{pos: Position{line:2,column:4}, value:"|"},
    }
	
	got, err := parse(input)
	
	if err != nil {
		t.Errorf("parse() returned %v ", err.Error())
	} else if !reflect.DeepEqual(want, got) {
		t.Errorf("parse() \nwant %+v, \ngot %+v ", want, got)
	}
	
}
