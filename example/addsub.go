package main

import (
	"fmt"
	"github.com/echotrue/timkit"
	"time"
)

func addSub() {
	tk := timkit.NewTimeKit(time.Now())
	fmt.Println(tk.AddYears(2))
	fmt.Println(tk.SubYears(2))

	fmt.Println(tk.AddQuarters(2))
	fmt.Println(tk.SubQuarters(2))

	fmt.Println(tk.AddMonths(2))
	fmt.Println(tk.SubMonths(2))

	fmt.Println(tk.AddDays(2))
	fmt.Println(tk.SubDays(2))
}