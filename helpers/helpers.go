package helpers

import (
	"log"
)

func Chk(err error) {
	if err != nil {
		log.Panic(err)
	}
}