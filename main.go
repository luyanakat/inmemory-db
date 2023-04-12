package main

import (
	"fmt"

	"github.com/luyanakat/inmemory-db/internal"
)

func main() {
	a := internal.GetID()

	fmt.Println(a)
}
