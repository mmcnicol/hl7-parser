package main

import (
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {

	//want := make([]Token, 0)
	//want = append(want, Token{Position{1,1}, "MSH"})
	
	want := []Token{
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
	
	/*
	input:=`MSH|^~\&|EPIC|EPICADT|SMS|SMSADT|199912271408|CHARRIS|ADT^A04|1817457|D|2.5|
PID||0493575^^^2^ID 1|454721||DOE^JOHN^^^^|DOE^JOHN^^^^|19480203|M||B|254 MYSTREET AVE^^MYTOWN^OH^44123^USA||(216)123-4567|||M|NON|400003403~1129086|
NK1||ROE^MARIE^^^^|SPO||(216)123-4567||EC|||||||||||||||||||||||||||
PV1||O|168 ~219~C~PMA^^^^^^^^^||||277^ALLEN MYLASTNAME^BONNIE^^^^|||||||||| ||2688684|||||||||||||||||||||||||199912271408||||||002376853
	`
	*/
	
	input:=`MSH|^~\&|EPIC|EPICADT|SMS|SMSADT|199912271408|CHARRIS|ADT^A04|
PID|`
	
	got, err := tokenizer(input)
	
	if err != nil {
		t.Errorf("parse() returned %v ", err.Error())
	} else if !reflect.DeepEqual(want, got) {
		t.Errorf("parse() \nwant %+v, \ngot %+v ", want, got)
	}
	
}
