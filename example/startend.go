package main

import (
	"fmt"
	"github.com/echotrue/timkit"
	"time"
)

func startEnd() {
	tk := timkit.NewTimeKit(time.Now())

	fmt.Println(tk.StartOfCentury())
	fmt.Println(tk.StartOfQuarter())
	fmt.Println(tk.StartOfYear())
	fmt.Println(tk.StartOfMonth())
	fmt.Println(tk.StartOfDay())
	fmt.Println(tk.StartOfWeek())
}
