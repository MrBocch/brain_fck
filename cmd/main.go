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
			case ',':
				continue
			case '#':
				// printing entire tape
				ptape(tape, p)


			default:
				continue
		}
	}
}

func ptape(t [size]uint8, p int32) {
	fmt.Printf("\n==DEBUG==\nPointer Addrres: %5X\n", p)
	for i, v := range t {
		if i % 8 == 0 {
			fmt.Println()
			fmt.Printf("%5X: ", i)
		}
		fmt.Printf("[%3v]", v)
	}
	fmt.Println("\n==DEBUG==")
}
