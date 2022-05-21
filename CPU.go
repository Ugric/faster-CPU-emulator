package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type CPU_type struct {
	load      func([]int64)
	run       func()
	PC        int64
	registers []int64
	memory    []int64
}

func CPU(registerSize int64, memorySize int64) CPU_type {
	var registers = make([]int64, registerSize)
	var memory = make([]int64, memorySize)
	var PC int64 = 0
	exec := func() {
		if memory[PC] == 1 {
			registers[memory[PC+1]] = registers[memory[PC+2]] + registers[memory[PC+3]]
			PC += 4
		} else if memory[PC] == 2 {
			registers[memory[PC+1]] = registers[memory[PC+2]] - registers[memory[PC+3]]
			PC += 4
		} else if memory[PC] == 3 {
			registers[memory[PC+1]] = registers[memory[PC+2]] * registers[memory[PC+3]]
			PC += 4
		} else if memory[PC] == 4 {
			registers[memory[PC+1]] = registers[memory[PC+2]] / registers[memory[PC+3]]
			PC += 4
		} else if memory[PC] == 5 {
			r := rune(registers[memory[PC+1]])
			fmt.Print(string(r))
			PC += 2
		} else if memory[PC] == 6 {
			// set memory to value of register
			memory[memory[PC+1]] = registers[memory[PC+2]]
			PC += 3
		} else if memory[PC] == 7 {
			//  set register
			registers[memory[PC+1]] = memory[PC+2]
			PC += 3
		} else if memory[PC] == 8 {
			//  set register from memory
			registers[memory[PC+1]] = memory[memory[PC+2]]
			PC += 3
		} else if memory[PC] == 9 {
			// sleep from register
			time.Sleep(time.Duration(registers[memory[PC+1]]) * time.Millisecond)
			PC += 2
		} else if memory[PC] == 10 {
			// jump if less than
			if registers[memory[PC+1]] < registers[memory[PC+2]] {
				PC = memory[PC+3]
			} else {
				PC += 4
			}
		} else if memory[PC] == 11 {
			// output number from register
			fmt.Print(registers[memory[PC+1]])
			PC += 2
		} else if memory[PC] == 12 {
			// set memory (index in a register) from a register
			memory[registers[memory[PC+1]]] = registers[memory[PC+2]]
			PC += 3
		}
	}
	run := func() {
		for memory[PC] != 0 {
			exec()
		}
	}
	load := func(script []int64) {
		for i := 0; i < len(script); i++ {
			memory[i] = script[i]
		}
	}
	return CPU_type{
		load:      load,
		run:       run,
		PC:        PC,
		registers: registers,
		memory:    memory,
	}
}

func main() {
	regsize, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	memorysize, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	my_CPU := CPU(regsize, memorysize)
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	my_list_of_strings := strings.Split(strings.ReplaceAll(string(content), "\n", " "), " ")
	my_list_of_ints := make([]int64, len(my_list_of_strings))
	for i := 0; i < len(my_list_of_strings); i++ {
		my_list_of_ints[i], _ = strconv.ParseInt(my_list_of_strings[i], 10, 64)
	}
	my_list_of_strings = nil
	my_CPU.load(my_list_of_ints)
	my_CPU.run()
}
