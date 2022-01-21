package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type customer struct {
	itemCount      int16
	typeOfCustomer string
	inService      bool
	timeArrived    int16
}

func cashier(x string) {
	file, err := os.Open("./" + x + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileLine := 0
	numberOfRegisters := 0

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
			numberOfRegisters = intVar
			fmt.Println(numberOfRegisters)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// var x customer = customer{ itemCount: 13, typeOfCustomer: "A", inService: true, timeArrived: 2 };

func (x customer) GetItemCount() {
	fmt.Println(x.itemCount)
}

// func (x customer)

func lineOfCustomer() {

}

func main() {
	var local string
	fmt.Println("qual o nome do arquivo?")
	fmt.Scan(&local)
	cashier(local)
}
