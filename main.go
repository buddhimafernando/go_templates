package main

import (
	"fmt"
	"os"
	"text/template"
)

type MonthlySales struct {
	Month  string
	PreviousRevenue  int
	MonthlyRevenue int
	Profit int
}

func main() {
	salesInfo := MonthlySales{"January", 900000, 1000000, 0}
	profit := salesInfo.PreviousRevenue - salesInfo.MonthlyRevenue
	salesInfo.Profit = profit

	tmpl, err := template.ParseFiles("template.txt")
	if err!= nil{
		fmt.Println(err)
	}

	err = tmpl.Execute(os.Stdout, salesInfo)
	if err!= nil{
		fmt.Println(err)
	}
}