package model

import "fmt"

type Student struct {
	ID     string
	Name   string
	Age    int
	School string
}

func (s Student) String() string {
	return fmt.Sprintf("ID: %s\nName: %s,\nAge: %d,\nSchool:%s\n\n", s.ID, s.Name, s.Age, s.School)
}
