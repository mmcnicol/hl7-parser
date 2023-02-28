package main

import (
	//"fmt"
	//"strings"
)

func preParse(tokens []Token) ([]Token, error) {
	
    //for _, token := range tokens {
	for i := 0; i < len(tokens); i++ {

		if tokens[i].level < 1 {
			tokens[i].level = 1
		}
		
		if tokens[i].pos.line == 1 && tokens[i].pos.column <= 8 { // TODO: for MSH only
			if tokens[i].value == "^" {
				tokens[i].level = 2
			} else if tokens[i].value == "~" {
				tokens[i].level = 2
			} else if tokens[i].value == "\\&" {
				tokens[i].level = 2
			}
		} else if tokens[i].value == "^" {
			if tokens[i-1].level < 3 {
				tokens[i-1].level = 2
			}
			tokens[i].level = 2
			tokens[i+1].level = 2
		} else if tokens[i].value == "~" {
			tokens[i-1].level = 3
			tokens[i].level = 3
			tokens[i+1].level = 3
		}
		
    }
	
	return tokens, nil
}
