package main

import (
	"fmt"
	"github.com/echotrue/timkit"
	"log"
	"time"
)

func main() {
	fmt.Println(timkit.Now())
	fmt.Println(timkit.Now().AddDays(12))
	fmt.Println(timkit.Now().SubDays(12))
	fmt.Println(timkit.Now().StartOfWeek())
	fmt.Println(timkit.Now().EndOfWeek())
	fmt.Println(timkit.Now().IsWeekday())

	l, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal(err)
	}
	ntk := timkit.NewTimeKit(time.Date(2021, 1, 2, 15, 4, 5, 0, l))
	fmt.Println(timkit.Now().DiffInDays(ntk, true))

	fmt.Println(timkit.Now().DiffInDaysFiltered(ntk, func(kit *timkit.TimeKit) bool {
		return kit.IsWeekday()
	}, true))
}
