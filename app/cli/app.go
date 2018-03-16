package main

import (
	"log"

	"github.com/Gujarats/ark"
)

func main() {
	sess, err := ark.CreateSession("ap-southeast-1")
	if err != nil {
		log.Fatal(err)
	}

	err = ark.StoreKeys(sess, "testing")
	if err != nil {
		log.Fatal(err)
	}

}
