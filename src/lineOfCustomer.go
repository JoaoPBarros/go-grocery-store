package lineOfCustomer

import "fmt"

type customer struct {
	itemCount      int16
	typeOfCustomer string
	inService      bool
	timeArrived    int16
}

// var x customer = customer{ itemCount: 13, typeOfCustomer: "A", inService: true, timeArrived: 2 };

func (x customer) GetItemCount() {
	fmt.Println(x.itemCount)
}

func (x customer)

func lineOfCustomer() {

}
