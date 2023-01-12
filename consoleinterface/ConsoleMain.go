package consoleinterface

import (
	"fmt"
	"os"
	_ "strconv"
)

func ConsoleMain() {

	fmt.Println("Choose the table of your choice to perform the task")
	fmt.Printf("1.Product_master\n2.Category_master\n3.Inventory\n4.CartItem\n5.Quit\n")
	fmt.Println("Enter choice")
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println(err)
	}

	if choice == 1 {
		ProductMaster()
	} else if choice == 2 {
		CategoryMaster()
	} else if choice == 3 {
		Inventory()
	} else if choice == 4 {
		Cart()
	} else if choice == 5 {
		os.Exit(0)
	}

}
