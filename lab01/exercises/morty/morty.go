package morty

import (
	"fmt"
)

func Rick(name string) {
	fmt.Println("Test from Morty package", name)
}

func ShowXTimes(X int, name string) {
	fmt.Printf("Hei %s, Get in the portal Rick, it's only a %d minutes adventure!\n", name, X)
}
