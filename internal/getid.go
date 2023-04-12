package internal

import (
	"log"

	uuid "github.com/nu7hatch/gouuid"
)

func GetID() string {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	return u.String()
}
