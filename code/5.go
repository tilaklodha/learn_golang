package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 10000,
	}

	gohan := Saiyan{}

	gohan.Name = "Gohan"
	gohan.Power = 9000

	vegeta := Saiyan{
		Name: "Vegeta",
	}
	vegeta.Power = 9000

	fmt.Println(goku)
	fmt.Println(gohan.Name)
	fmt.Println(vegeta.Power)
}
