package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type customer struct {
	itemCount      int8
	typeOfCustomer string
	inService      bool
	timeArrived    int8
}

type register struct {
	position      int
	training      bool
	arrOfCustomer []customer
}

var registers = make([]register, 0)

func createQtOfRegisters(x int) {
	registers := make([]register, x)

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

			fmt.Println(scanner.Text())

		} else {
			fileLine = fileLine + 1

			intVar, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

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
