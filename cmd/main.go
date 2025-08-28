package main

import (
	"fmt"
	"os"
)

const size  = 30_000
var tape [size] uint8
var p int32 = 0

func main(){
	if len(os.Args) == 1 {
		fmt.Println("REPL mode?")
	} else {
		Run(os.Args[1])
	}
}

func Run(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error opening %v\n", filePath)
		os.Exit(1)
	}
	var p int32 = 0
	var code = string(data)

	for _, char := range(code){
		switch char {
			case '+':
				tape[p] += 1
			case '-':
				tape[p] -= 1
			case '>':
				p += 1
			case '<':
				p -= 1
			case '.':
				fmt.Printf("%c", tape[p])
			// printing entire tape
			case ':':
				ptape(tape)


			default:
				continue
		}
	}
}

func ptape(t [size]uint8) {
	for _, v := range t {
		fmt.Println(v)
	}
}
