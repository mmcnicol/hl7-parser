package main

import (
	"reflect"
	"testing"
)

func TestPreParserMultipleLevelOneComposites(t *testing.T) {

	want := []Token{
        {pos: Position{line:1,column:3}, value:"MSH", level: 1},
        {pos: Position{line:1,column:4}, value:"|", level: 1},
        {pos: Position{line:1,column:5}, value:"^", level: 2},
        {pos: Position{line:1,column:6}, value:"~", level: 2},
		{pos: Position{line:1,column:8}, value:"\\&", level: 2},
		{pos: Position{line:1,column:9}, value:"|", level: 1},
		{pos: Position{line:1,column:12}, value:"AAA", level: 1},
		{pos: Position{line:1,column:13}, value:"|", level: 1},
		{pos: Position{line:1,column:16}, value:"BBB", level: 1},
		{pos: Position{line:1,column:17}, value:"|", level: 1},
		{pos: Position{line:1,column:20}, value:"CCC", level: 1},
		{pos: Position{line:1,column:21}, value:"|", level: 1},
		{pos: Position{line:1,column:24}, value:"DDD", level: 1},
		{pos: Position{line:1,column:25}, value:"|", level: 1},
    }
		
	input := []Token{
        {pos: Position{line:1,column:3}, value:"MSH"},
        {pos: Position{line:1,column:4}, value:"|"},
        {pos: Position{line:1,column:5}, value:"^"},
        {pos: Position{line:1,column:6}, value:"~"},
		{pos: Position{line:1,column:8}, value:"\\&"},
		{pos: Position{line:1,column:9}, value:"|"},
		{pos: Position{line:1,column:12}, value:"AAA"},
		{pos: Position{line:1,column:13}, value:"|"},
		{pos: Position{line:1,column:16}, value:"BBB"},
		{pos: Position{line:1,column:17}, value:"|"},
		{pos: Position{line:1,column:20}, value:"CCC"},
		{pos: Position{line:1,column:21}, value:"|"},
		{pos: Position{line:1,column:24}, value:"DDD"},
		{pos: Position{line:1,column:25}, value:"|"},
    }
	
	got, err := preParse(input)
	
	if err != nil {
		t.Errorf("parse() returned %v ", err.Error())
	} else if !reflect.DeepEqual(want, got) {
		t.Errorf("parse() \nwant %+v \ngot  %+v ", want, got)
	}
}

func TestPreParserMultipleLevelTwoComposites(t *testing.T) {

	want := []Token{
        {pos: Position{line:1,column:3}, value:"MSH", level: 1},
        {pos: Position{line:1,column:4}, value:"|", level: 1},
        {pos: Position{line:1,column:5}, value:"^", level: 2},
        {pos: Position{line:1,column:6}, value:"~", level: 2},
		{pos: Position{line:1,column:8}, value:"\\&", level: 2},
		{pos: Position{line:1,column:9}, value:"|", level: 1},
		{pos: Position{line:1,column:12}, value:"AAA", level: 2},
		{pos: Position{line:1,column:13}, value:"^", level: 2},
		{pos: Position{line:1,column:16}, value:"BBB", level: 2},
		{pos: Position{line:1,column:17}, value:"^", level: 2},
		{pos: Position{line:1,column:20}, value:"CCC", level: 2},
		{pos: Position{line:1,column:21}, value:"^", level: 2},
		{pos: Position{line:1,column:24}, value:"DDD", level: 2},
		{pos: Position{line:1,column:25}, value:"|", level: 1},
    }
		
	input := []Token{
        {pos: Position{line:1,column:3}, value:"MSH"},
        {pos: Position{line:1,column:4}, value:"|"},
        {pos: Position{line:1,column:5}, value:"^"},
        {pos: Position{line:1,column:6}, value:"~"},
		{pos: Position{line:1,column:8}, value:"\\&"},
		{pos: Position{line:1,column:9}, value:"|"},
		{pos: Position{line:1,column:12}, value:"AAA"},
		{pos: Position{line:1,column:13}, value:"^"},
		{pos: Position{line:1,column:16}, value:"BBB"},
		{pos: Position{line:1,column:17}, value:"^"},
		{pos: Position{line:1,column:20}, value:"CCC"},
		{pos: Position{line:1,column:21}, value:"^"},
		{pos: Position{line:1,column:24}, value:"DDD"},
		{pos: Position{line:1,column:25}, value:"|"},
    }
	
	got, err := preParse(input)
	
	if err != nil {
		t.Errorf("parse() returned %v ", err.Error())
	} else if !reflect.DeepEqual(want, got) {
		t.Errorf("parse() \nwant %+v \ngot  %+v ", want, got)
	}
}

func TestPreParserMultipleLevelThreeComposites(t *testing.T) {

	want := []Token{
        {pos: Position{line:1,column:3}, value:"MSH", level: 1},
        {pos: Position{line:1,column:4}, value:"|", level: 1},
        {pos: Position{line:1,column:5}, value:"^", level: 2},
        {pos: Position{line:1,column:6}, value:"~", level: 2},
		{pos: Position{line:1,column:8}, value:"\\&", level: 2},
		{pos: Position{line:1,column:9}, value:"|", level: 1},
		{pos: Position{line:1,column:12}, value:"AAA", level: 2},
		{pos: Position{line:1,column:13}, value:"^", level: 2},
		{pos: Position{line:1,column:16}, value:"BBB", level: 3},
		{pos: Position{line:1,column:17}, value:"~", level: 3},
		{pos: Position{line:1,column:20}, value:"CCC", level: 3},
		{pos: Position{line:1,column:21}, value:"~", level: 3},
		{pos: Position{line:1,column:24}, value:"DDD", level: 3},
		{pos: Position{line:1,column:25}, value:"^", level: 2},
		{pos: Position{line:1,column:28}, value:"EEE", level: 2},
		{pos: Position{line:1,column:29}, value:"|", level: 1},
    }
		
	input := []Token{
        {pos: Position{line:1,column:3}, value:"MSH"},
        {pos: Position{line:1,column:4}, value:"|"},
        {pos: Position{line:1,column:5}, value:"^"},
        {pos: Position{line:1,column:6}, value:"~"},
		{pos: Position{line:1,column:8}, value:"\\&"},
		{pos: Position{line:1,column:9}, value:"|"},
		{pos: Position{line:1,column:12}, value:"AAA"},
		{pos: Position{line:1,column:13}, value:"^"},
		{pos: Position{line:1,column:16}, value:"BBB"},
		{pos: Position{line:1,column:17}, value:"~"},
		{pos: Position{line:1,column:20}, value:"CCC"},
		{pos: Position{line:1,column:21}, value:"~"},
		{pos: Position{line:1,column:24}, value:"DDD"},
		{pos: Position{line:1,column:25}, value:"^"},
		{pos: Position{line:1,column:28}, value:"EEE"},
		{pos: Position{line:1,column:29}, value:"|"},
    }
	
	got, err := preParse(input)
	
	if err != nil {
		t.Errorf("parse() returned %v ", err.Error())
	} else if !reflect.DeepEqual(want, got) {
		t.Errorf("parse() \nwant %+v \ngot  %+v ", want, got)
	}
}

func TestPreParser(t *testing.T) {

	want := []Token{
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
	
	got, err := preParse(input)
	
	if err != nil {
		t.Errorf("parse() returned %v ", err.Error())
	} else if !reflect.DeepEqual(want, got) {
		t.Errorf("parse() \nwant %+v \ngot  %+v ", want, got)
	}
}
