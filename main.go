package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	fmt.Println("Created the bill - ", b.name)
	return b
}

func getInput(promrpt string, reader *bufio.Reader) (string, error) {
	fmt.Print(promrpt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promtOptions(b)
		}

		b.addItems(name, p)
		fmt.Println("Item added: ", name, price)
		promtOptions(b)
		break
	case "t":
		tip, _ := getInput("Enter tip: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promtOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip added: ", tip)
		promtOptions(b)
		break
	case "s":
		fmt.Println("you chose to save the bill", b)
		b.saveBill()
		break
	default:
		fmt.Println("invalid option")
		promtOptions(b)
	}
}

func main() {
	myBill := createBill()
	promtOptions(myBill)
}
