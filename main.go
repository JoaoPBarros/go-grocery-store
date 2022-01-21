package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type customer struct {
	typeOfCustomer string
	timeArrived    int
	itemCount      int
	inService      bool
}

type register struct {
	position      int
	training      bool
	arrOfCustomer []customer
}

var registers = make([]register, 0)

func createQtOfRegisters(x int) {
	registers := make([]register, (x + 1))

	if x != 1 {
		for i := 0; i < (x - 1); i++ {

			registers[i] = register{
				position: i + 1,
				training: false,
			}
		}
	}

	registers[x] = register{
		position: x,
		training: true,
	}
}

var t int = 0
var registersQty int = 0

func resolveClientPosition(x [3]string) {
	for i := 0; i < len(registers); i++ {
		if len(registers[i].arrOfCustomer) == 0 {
			intVar, err := strconv.Atoi(x[1])
			if err != nil {
				log.Fatal(err)
			}

			intVar2, err2 := strconv.Atoi(x[2])
			if err2 != nil {
				log.Fatal(err2)
			}

			registers[i].arrOfCustomer = append(registers[i].arrOfCustomer, customer{
				typeOfCustomer: x[0],
				timeArrived:    intVar,
				itemCount:      intVar2,
			})
		}
		registersQty = registersQty + 1
	}
}

// function to read each file line in root of the project

func cashier(x string) {
	file, err := os.Open("./" + x + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileLine := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if fileLine != 0 {
			fileLine = fileLine + 1

			// fmt.Println(scanner.Text())
			fmt.Println(strings.Split(scanner.Text(), " "))

		}
		if fileLine == 0 {
			fileLine = fileLine + 1

			intVar, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(intVar)
			createQtOfRegisters(intVar)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// var x customer = customer{ itemCount: 13, typeOfCustomer: "A", inService: true, timeArrived: 2 };

// func (x customer) GetItemCount() {
// 	fmt.Println(x.itemCount)
// }

func main() {
	var local string

	fmt.Println("qual o nome do arquivo?")
	fmt.Scan(&local)
	cashier(local)

}
