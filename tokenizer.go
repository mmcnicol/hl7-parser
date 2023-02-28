package main

import (
	"fmt"
	//"strings"
)

type Position struct {
	line   int
	column int
}

type Token struct {
	pos Position
	value string
}

func tokenizer(message string) ([]Token, error) {
	fmt.Println("hello!")
	
	tokens := make([]Token, 0)
	
	//var tokens []Token{}
	line := 1
	value := ""
	column := 1

    for _, runeValue := range message {

		pos := Position{}
		pos.line = line
		
        switch(runeValue) {
			case '|':
				if len(value) > 0  {
					pos.column = column - 1
					token := Token{}
					token.pos = pos
					token.value = string(value)
					tokens = append(tokens, token)
				}
				pos.column = column
				token := Token{}
				token.pos = pos
				token.value = string(runeValue)
				tokens = append(tokens, token)
				value = ""
			case '^':
				if len(value) > 0  {
					pos.column = column - 1
					token := Token{}
					token.pos = pos
					token.value = string(value)
					tokens = append(tokens, token)
				}
				pos.column = column
				token := Token{}
				token.pos = pos
				token.value = string(runeValue)
				tokens = append(tokens, token)
				value = ""
			case '~':
				if len(value) > 0  {
					pos.column = column - 1
					token := Token{}
					token.pos = pos
					token.value = string(value)
					tokens = append(tokens, token)
				}
				pos.column = column
				token := Token{}
				token.pos = pos
				token.value = string(runeValue)
				tokens = append(tokens, token)
				value = ""
			case '\n':
				if len(value) > 0  {
					pos.column = column - 1
					token := Token{}
					token.pos = pos
					token.value = string(value)
					tokens = append(tokens, token)
				}
				pos.column = column
				token := Token{}
				token.pos = pos
				token.value = "newline"
				tokens = append(tokens, token)
				line = line +1
				value = ""
				column = 0
			case '\r':
				pos.column = column
				token := Token{}
				token.pos = pos
				token.value = "carriage return"
				tokens = append(tokens, token)
				value = ""
				column = 0
			default:
				//fmt.Printf("%c\n", runeValue)
				value = value + string(runeValue)
		}
		column = column + 1
    }
	
	return tokens, nil
}
