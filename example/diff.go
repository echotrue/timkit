package main

import (
	"fmt"
	"github.com/echotrue/timkit"
	"log"
	"time"
)

func diff() {
	tk := timkit.NewTimeKit(time.Now())

	l, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal(err)
	}
	ntk := timkit.NewTimeKit(time.Date(2021, 1, 2, 15, 4, 5, 0, l))

	fmt.Println(tk.DiffInMonths(ntk, true))
	fmt.Println(tk.DiffInDays(ntk, true))
	fmt.Println(tk.DiffInHours(ntk, true))
	fmt.Println(tk.DiffInMinutes(ntk, true))
	fmt.Println(tk.DiffInSeconds(ntk, true))

	fmt.Println(tk.DiffInDaysFiltered(ntk, func(kit *timkit.TimeKit) bool {
		return kit.IsWeekday()
	}, true))
}
