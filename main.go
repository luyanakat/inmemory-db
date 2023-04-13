package main

import (
	"fmt"
	"log"

	"github.com/luyanakat/inmemory-db/cmd"
	"github.com/luyanakat/inmemory-db/internal"
	"github.com/luyanakat/inmemory-db/model"
)

func main() {
	s, err := cmd.NewStore(20)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Free()

	err = s.AddList(19)
	if err != nil {
		log.Fatal(err)
	}

	student3 := model.Student{
		ID:     "123",
		Name:   "TCC",
		Age:    21,
		School: "KMA",
	}
	err = s.AddPeople(&student3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Print Asc Order")
	s.PrintAsOrder(internal.Desc)

	fmt.Println("List By Age:")
	s.PrintByAge(12)
}
