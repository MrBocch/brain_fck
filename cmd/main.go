package main

import (
	"fmt"
	"os"
)

const SIZE = 30_000
var tape [SIZE] uint8

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

		case '[':
			if tape[dp] == 0 {
				ip = skips(code, ip)
				continue
			} else {
				ip += 1
			}
		case ']': {
			if tape[dp] != 0 {
				ip = skips_back(code, ip)
				continue
			} else {
				ip += 1
			}

		}


		default:
			ip += 1
			continue
		}
	}
}

func ptape(t [SIZE]uint8, dp int32) {
	fmt.Printf("\n==DEBUG==\nPointer Addrres: 0x%04X\n", dp)
	for i, v := range t {
		if i % 8 == 0 {
			fmt.Println()
			fmt.Printf("0x%04X: ", i)
		}
		fmt.Printf("[%03v]", v)
	}
	fmt.Println("\n==DEBUG==")
}

func skips(code string, ip int) int {
	ip += 1
	lvl := 0
	for ;; {
		c := code[ip]
		if c == '[' {
			lvl += 1
			ip += 1
			continue
		}

		if c != ']' {
			ip += 1
		}
		if c == ']' {
			if lvl != 0 {
				lvl -= 1
				ip += 1
			} else {
				ip += 1
				return ip
			}
		}
	}
}

func skips_back(code string, ip int) int {
	ip -= 1
	lvl := 0
	for ;; {
		c := code[ip]
		if c == ']' {
			lvl += 1
			ip -= 1
			continue
		}

		if c != '[' {
			ip -= 1
		}
		if c == '[' {
			if lvl != 0 {
				lvl -= 1
				ip -= 1
			} else {
				// cant just copy and paste
				return ip
			}
		}
	}
}
