package cashier

import "fmt"

func Cashier() {
	fmt.Println("teste")
}

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func cashier(x string) {
// 	file, err := os.Open(x)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	// optionally, resize scanner's capacity for lines over 64K, see next example
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
