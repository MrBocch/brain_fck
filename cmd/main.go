package main

import (
	"fmt"
	"os"
)

const size  = 30_000
var tape [size] uint8

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

	var dp int32 = 0
	var ip int = 0
	var code = string(data)

	for ;; {
		if ip == len(code) { break }
		char := code[ip]
		switch char {
		case '+':
			tape[dp] += 1
			ip += 1
		case '-':
			tape[dp] -= 1
			ip += 1
		case '>':
			dp += 1
			ip += 1
		case '<':
			dp -= 1
			ip += 1
		case '.':
			fmt.Printf("%c", tape[dp])
			ip += 1
		case ',':
			continue
		case '#':
			ptape(tape, dp)
			ip += 1
		default:
			ip += 1
			continue
		}
	}
}

func ptape(t [size]uint8, p int32) {
	fmt.Printf("\n==DEBUG==\nPointer Addrres: 0x%04X\n", p)
	for i, v := range t {
		if i % 8 == 0 {
			fmt.Println()
			fmt.Printf("0x%04X: ", i)
		}
		fmt.Printf("[%03v]", v)
	}
	fmt.Println("\n==DEBUG==")
}
