package main

import (
	"fmt"
	"log"
	convert "roman-numbers/convertions"
)

func main() {

	var number string
	fmt.Println("Enter the roman number to convert:")
	fmt.Scanf("%s", &number)

	result, err := convert.ConvertRomanToNumber(number)

	if err != nil {
		fmt.Println("Something went wrong")
		log.Fatal(err)
	}

	fmt.Println("---------------The result is ------------------")

	fmt.Println(result)
}
