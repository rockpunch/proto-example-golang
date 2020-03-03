package main

import (
	"fmt"
	"github.com/rockpunch/proto-example-golang/example-3/proto"
)

func main() {
	doEnum()
}


func doEnum() {
	em := proto.EnumMessage{
		Id: 42,
		DayOfTheWeek: proto.DayOfTheWeek_MONDAY,
	}

	em.DayOfTheWeek = proto.DayOfTheWeek_FRIDAY

	fmt.Println(em)
}