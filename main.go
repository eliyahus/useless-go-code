package main

import (
	"fmt"
	"github.com/eliyahus/useless-go-code/eatingunicorns"
)

func main() {
	u := eatingunicorns.NewUnicorn("Glory")
	fmt.Println(u)
	fmt.Printf("Is %s eatable? %t", u.Name(), eatingunicorns.Eatable(&u))
}
